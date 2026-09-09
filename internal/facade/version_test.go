package facade

import (
	"strings"
	"testing"
)

func TestNewVersionString(t *testing.T) {
	got := NewVersionString("v1.2.3").String()
	want := "today v1.2.3\nhttps://github.com/goark/today"
	if got != want {
		t.Fatalf("NewVersionString().String() = %q, want %q", got, want)
	}
}

func TestUsageString(t *testing.T) {
	got := usageString()
	want := "today\n\nUsage:\n  today [flags] [yyyy-mm-dd]\n\nFlags"
	if got != want {
		t.Fatalf("usageString() = %q, want %q", got, want)
	}
}

func TestReplaceVersion_WithExplicitVersion(t *testing.T) {
	got := replaceVersion("v9.9.9")
	if got != "v9.9.9" {
		t.Fatalf("replaceVersion() = %q, want %q", got, "v9.9.9")
	}
}

func TestNewVersionString_EmptyVersionUsesFallback(t *testing.T) {
	got := NewVersionString("").String()
	lines := strings.Split(got, "\n")
	if len(lines) != 2 {
		t.Fatalf("version string line count = %d, want %d: %q", len(lines), 2, got)
	}
	if !strings.HasPrefix(lines[0], "today ") {
		t.Fatalf("first line = %q, want prefix %q", lines[0], "today ")
	}
	if lines[0] == "today " {
		t.Fatal("first line has empty version payload")
	}
	if lines[1] != "https://github.com/goark/today" {
		t.Fatalf("second line = %q, want %q", lines[1], "https://github.com/goark/today")
	}
}

func TestJoinNonEmpty(t *testing.T) {
	got := joinNonEmpty("abc123", "", "(compiled with go1.26.5)")
	want := "abc123 (compiled with go1.26.5)"
	if got != want {
		t.Fatalf("joinNonEmpty() = %q, want %q", got, want)
	}
}

/* Copyright 2026 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
