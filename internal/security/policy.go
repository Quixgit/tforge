package security

import "fmt"

type Policy struct {
	AllowApply   bool
	AllowDestroy bool
}

func NewPolicy(allowApply, allowDestroy bool) Policy {
	return Policy{
		AllowApply:   allowApply,
		AllowDestroy: allowDestroy,
	}
}

func (p Policy) Check(action string) error {
	switch action {
	case "destroy":
		if !p.AllowDestroy {
			return fmt.Errorf("destroy is disabled by safety policy")
		}
	case "apply":
		if !p.AllowApply {
			return fmt.Errorf("apply is disabled by safety policy")
		}
	}

	return nil
}
