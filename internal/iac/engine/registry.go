package engine

import (
	"errors"
)

type Registry struct {
	engines map[string]Engine
}

func NewRegistry() Registry {
	return Registry{
		engines: map[string]Engine{},
	}
}

func (r *Registry) Register(e Engine) {
	r.engines[e.Name()] = e
}

func (r Registry) Get(name string) (Engine, error) {
	e, ok := r.engines[name]
	if !ok {
		return nil, errors.New("engine not registered: " + name)
	}

	return e, nil
}

func (r Registry) Names() []string {
	names := make([]string, 0, len(r.engines))
	for name := range r.engines {
		names = append(names, name)
	}

	return names
}
