package security

import (
	"strings"
	"testing"
)

func TestMaskLine(t *testing.T) {
	line := `password = "super-secret"`
	masked := MaskLine(line)

	if strings.Contains(masked, "super-secret") {
		t.Fatalf("secret was not masked: %s", masked)
	}

	if !strings.Contains(masked, "<masked>") {
		t.Fatalf("expected masked marker")
	}
}

func TestMaskToken(t *testing.T) {
	line := `token = abc123`
	masked := MaskLine(line)

	if strings.Contains(masked, "abc123") {
		t.Fatalf("token was not masked")
	}
}
