package execution

import (
	"regexp"
	"strings"
	"time"
)

var (
	startRe = regexp.MustCompile(`^([^\s]+): (Creating|Modifying|Destroying)\.\.\.$`)
	doneRe  = regexp.MustCompile(`^([^\s]+): (Creation|Modifications|Destruction) complete`)
	failRe  = regexp.MustCompile(`^Error: `)
)

type EventType string

const (
	EventStart  EventType = "start"
	EventDone   EventType = "done"
	EventError  EventType = "error"
	EventIgnore EventType = "ignore"
)

type Event struct {
	Type      EventType
	Address   string
	Action    string
	Error     string
	Timestamp time.Time
}

func ParseLine(line string) Event {
	line = strings.TrimSpace(line)

	if line == "" {
		return Event{Type: EventIgnore}
	}

	if failRe.MatchString(line) {
		return Event{
			Type:      EventError,
			Error:     line,
			Timestamp: time.Now(),
		}
	}

	if m := startRe.FindStringSubmatch(line); len(m) > 0 {
		return Event{
			Type:      EventStart,
			Address:   m[1],
			Action:    strings.ToLower(m[2]),
			Timestamp: time.Now(),
		}
	}

	if m := doneRe.FindStringSubmatch(line); len(m) > 0 {
		return Event{
			Type:      EventDone,
			Address:   m[1],
			Timestamp: time.Now(),
		}
	}

	return Event{
		Type: EventIgnore,
	}
}
