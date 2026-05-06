package app

import "fmt"

func errUnsupportedAction(action string) error {
	return fmt.Errorf("unsupported action: %s", action)
}
