package execution

import "testing"

func TestParseStart(t *testing.T) {
	ev := ParseLine("aws_vpc.main: Creating...")

	if ev.Type != EventStart {
		t.Fatalf("expected start")
	}

	if ev.Address != "aws_vpc.main" {
		t.Fatalf("bad address")
	}
}

func TestParseDone(t *testing.T) {
	ev := ParseLine("aws_vpc.main: Creation complete after 3s")

	if ev.Type != EventDone {
		t.Fatalf("expected done")
	}
}

func TestParseError(t *testing.T) {
	ev := ParseLine("Error: something exploded")

	if ev.Type != EventError {
		t.Fatalf("expected error")
	}
}
