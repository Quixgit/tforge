package events

import "time"

type Type string

const (
	TypeStarted  Type = "started"
	TypeStdout   Type = "stdout"
	TypeStderr   Type = "stderr"
	TypeFinished Type = "finished"
	TypeError    Type = "error"
)

type Event struct {
	Type      Type
	Engine    string
	Command   string
	Line      string
	Error     string
	ExitCode  int
	Timestamp time.Time
}

func New(t Type, engine, command string) Event {
	return Event{
		Type:      t,
		Engine:    engine,
		Command:   command,
		Timestamp: time.Now(),
	}
}
