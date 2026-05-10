package app

import "os/exec"

type SecurityTools struct {
	TFLint bool
	TFSec  bool
}

func detectSecurityTools() SecurityTools {
	return SecurityTools{
		TFLint: hasBinary("tflint"),
		TFSec:  hasBinary("tfsec"),
	}
}

func hasBinary(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
