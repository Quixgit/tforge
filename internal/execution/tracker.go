package execution

import "time"

type Tracker struct {
	resources map[string]*ResourceState
}

func NewTracker() *Tracker {
	return &Tracker{
		resources: map[string]*ResourceState{},
	}
}

func (t *Tracker) Reset() {
	t.resources = map[string]*ResourceState{}
}

func (t *Tracker) Handle(event Event) {
	switch event.Type {
	case EventStart:
		t.resources[event.Address] = &ResourceState{
			Address:   event.Address,
			Action:    event.Action,
			Status:    StatusRunning,
			StartedAt: event.Timestamp,
		}

	case EventDone:
		if r, ok := t.resources[event.Address]; ok {
			r.Status = StatusComplete
			r.EndedAt = event.Timestamp
		}

	case EventError:
		now := time.Now()

		t.resources["__global__error"] = &ResourceState{
			Address:   "__global__error",
			Status:    StatusFailed,
			Error:     event.Error,
			StartedAt: now,
			EndedAt:   now,
		}
	}
}

func (t *Tracker) List() []ResourceState {
	out := make([]ResourceState, 0, len(t.resources))

	for _, r := range t.resources {
		out = append(out, *r)
	}

	return out
}
