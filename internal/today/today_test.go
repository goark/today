package today

import (
	"testing"
	"time"
)

func TestString(t *testing.T) {
	got := String(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.UTC))
	want := "2026-09-09"
	if got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}
}
