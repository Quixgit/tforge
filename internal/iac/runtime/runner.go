package runtime

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/quix/tforge/internal/core/events"
)

var ErrUnsafeWorkingDir = errors.New("unsafe working directory")

type Runner struct {
	Timeout time.Duration
}

type CommandSpec struct {
	Engine  string
	Binary  string
	Dir     string
	Command string
	Args    []string
	Env     []string
}

func NewRunner(timeout time.Duration) Runner {
	if timeout <= 0 {
		timeout = 30 * time.Minute
	}

	return Runner{Timeout: timeout}
}

func (r Runner) Stream(ctx context.Context, spec CommandSpec) (<-chan events.Event, error) {
	if err := validateSpec(spec); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, r.Timeout)

	out := make(chan events.Event, 64)

	cmd := exec.CommandContext(ctx, spec.Binary, spec.Args...)
	cmd.Dir = spec.Dir
	cmd.Env = safeEnv(spec.Env)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		cancel()
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		cancel()
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		cancel()
		return nil, err
	}

	go func() {
		defer close(out)
		defer cancel()

		out <- events.New(events.TypeStarted, spec.Engine, spec.Command)

		done := make(chan struct{})

		go scanPipe(stdout, out, spec, events.TypeStdout)
		go scanPipe(stderr, out, spec, events.TypeStderr)

		go func() {
			_ = cmd.Wait()
			close(done)
		}()

		select {
		case <-ctx.Done():
			ev := events.New(events.TypeError, spec.Engine, spec.Command)
			ev.Error = ctx.Err().Error()
			out <- ev

		case <-done:
			exitCode := cmd.ProcessState.ExitCode()

			ev := events.New(events.TypeFinished, spec.Engine, spec.Command)
			ev.ExitCode = exitCode
			out <- ev
		}
	}()

	return out, nil
}

func (r Runner) Output(ctx context.Context, spec CommandSpec) ([]byte, error) {
	if err := validateSpec(spec); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, r.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, spec.Binary, spec.Args...)
	cmd.Dir = spec.Dir
	cmd.Env = safeEnv(spec.Env)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return stdout.Bytes(), fmt.Errorf("%w: %s", err, maskSecrets(stderr.String()))
	}

	return stdout.Bytes(), nil
}

func scanPipe(pipe any, out chan<- events.Event, spec CommandSpec, typ events.Type) {
	scanner, ok := pipe.(interface {
		Read([]byte) (int, error)
	})
	if !ok {
		return
	}

	s := bufio.NewScanner(scanner)
	for s.Scan() {
		ev := events.New(typ, spec.Engine, spec.Command)
		ev.Line = maskSecrets(s.Text())
		out <- ev
	}
}

func validateSpec(spec CommandSpec) error {
	if spec.Binary == "" {
		return errors.New("binary is required")
	}

	if strings.Contains(spec.Binary, " ") {
		return errors.New("binary must not contain spaces")
	}

	if spec.Dir == "" {
		return errors.New("working directory is required")
	}

	clean, err := filepath.Abs(spec.Dir)
	if err != nil {
		return err
	}

	info, err := os.Stat(clean)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return ErrUnsafeWorkingDir
	}

	if clean == "/" {
		return ErrUnsafeWorkingDir
	}

	return nil
}

func safeEnv(extra []string) []string {
	base := os.Environ()
	return append(base, extra...)
}

func maskSecrets(s string) string {
	keys := []string{
		"AWS_SECRET_ACCESS_KEY",
		"AWS_SESSION_TOKEN",
		"GOOGLE_CREDENTIALS",
		"ARM_CLIENT_SECRET",
		"TF_VAR_password",
		"TF_VAR_token",
		"TF_TOKEN",
	}

	for _, key := range keys {
		if strings.Contains(s, key) {
			s = strings.ReplaceAll(s, key, "[masked]")
		}
	}

	return s
}
