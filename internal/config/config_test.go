package config

import (
	"bytes"
	"testing"
)

func TestParseVersion(t *testing.T) {
	cfg, err := Parse([]string{"--version"}, &bytes.Buffer{})
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if !cfg.ShowVersion {
		t.Fatal("Parse() ShowVersion = false, want true")
	}
}
