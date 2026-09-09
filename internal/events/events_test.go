package events

import (
	"strings"
	"testing"
	"time"

	"github.com/goark/koyomi/value"
)

func TestImportEvents(t *testing.T) {
	jsonInput := `[{"start":"2026-09-01","end":"2026-09-10","title":"event-a"}]`
	got, err := ImportEvents(strings.NewReader(jsonInput))
	if err != nil {
		t.Fatalf("ImportEvents() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("len(ImportEvents()) = %d, want %d", len(got), 1)
	}
	if got[0].Title != "event-a" {
		t.Fatalf("ImportEvents()[0].Title = %q, want %q", got[0].Title, "event-a")
	}
}

func TestImportEvents_InvalidJSON(t *testing.T) {
	_, err := ImportEvents(strings.NewReader(`[{"start":"2026-09-01"`))
	if err == nil {
		t.Fatal("ImportEvents() error = nil, want non-nil")
	}
}

func TestEventContain(t *testing.T) {
	e := Event{Start: "2026-09-01", End: "2026-09-10", Title: "event-a"}
	dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
	got, err := e.Contain(dt)
	if err != nil {
		t.Fatalf("Contain() error = %v", err)
	}
	if !got {
		t.Fatal("Contain() = false, want true")
	}
}

func TestEventContain_InvalidStart(t *testing.T) {
	e := Event{Start: "bad", End: "2026-09-10", Title: "event-a"}
	dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
	_, err := e.Contain(dt)
	if err == nil {
		t.Fatal("Contain() error = nil, want non-nil")
	}
}

func TestGetEventsContaining(t *testing.T) {
	dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
	events := []Event{
		{Start: "2026-09-01", End: "2026-09-10", Title: "event-a"},
		{Start: "2026-10-01", End: "2026-10-05", Title: "event-b"},
	}
	got, err := GetEventsContaining(dt, events)
	if err != nil {
		t.Fatalf("GetEventsContaining() error = %v", err)
	}
	if len(got) != 1 || got[0] != "event-a" {
		t.Fatalf("GetEventsContaining() = %v, want [event-a]", got)
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
