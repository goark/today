package today

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/goark/errs"
	"github.com/goark/koyomi/value"
	"github.com/goark/today/internal/ecode"
)

func TestNewInfoDateAndStrings(t *testing.T) {
	dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
	i := NewInfoDate(dt)
	if i == nil {
		t.Fatal("NewInfoDate() = nil, want non-nil")
	}
	if i.Year != 2026 || i.Month != 9 || i.Day != 9 {
		t.Fatalf("date fields = %04d-%02d-%02d, want 2026-09-09", i.Year, i.Month, i.Day)
	}
	ss := i.Strings()
	if len(ss) < 2 {
		t.Fatalf("len(Strings()) = %d, want >= 2", len(ss))
	}
	if !strings.HasPrefix(ss[0], "2026-09-09 ") {
		t.Fatalf("Strings()[0] = %q, want prefix %q", ss[0], "2026-09-09 ")
	}
}

func TestInfoDate_ImportEvents(t *testing.T) {
	dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
	i := NewInfoDate(dt)
	jsonInput := `[{"start":"2026-09-01","end":"2026-09-10","title":"event-a"}]`
	if err := i.ImportEvents(strings.NewReader(jsonInput)); err != nil {
		t.Fatalf("ImportEvents() error = %v", err)
	}
	if len(i.Events) != 1 || i.Events[0] != "event-a" {
		t.Fatalf("Events = %v, want [event-a]", i.Events)
	}
}

func TestInfoDate_Strings_NilReceiver(t *testing.T) {
	var i *InfoDate
	if got := i.Strings(); len(got) != 0 {
		t.Fatalf("nil.Strings() = %v, want []", got)
	}
}

func TestInfoDate_ImportHoliday(t *testing.T) {
	orig := holidayEventTitlesFetcher
	t.Cleanup(func() { holidayEventTitlesFetcher = orig })

	t.Run("success", func(t *testing.T) {
		holidayEventTitlesFetcher = func(dt value.DateJp, tempDir string) ([]string, error) {
			return []string{"holiday-a"}, nil
		}
		dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
		i := NewInfoDate(dt)
		if err := i.ImportHoliday("tempDir"); err != nil {
			t.Fatalf("ImportHoliday() error = %v", err)
		}
		if len(i.Events) != 1 || i.Events[0] != "holiday-a" {
			t.Fatalf("Events = %v, want [holiday-a]", i.Events)
		}
	})

	t.Run("error", func(t *testing.T) {
		holidayEventTitlesFetcher = func(dt value.DateJp, tempDir string) ([]string, error) {
			return nil, errors.New("holiday fetch error")
		}
		dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
		i := NewInfoDate(dt)
		err := i.ImportHoliday("tempDir")
		if err == nil {
			t.Fatal("ImportHoliday() error = nil, want non-nil")
		}
		if !strings.Contains(err.Error(), "holiday fetch error") {
			t.Fatalf("ImportHoliday() error = %v, want contains %q", err, "holiday fetch error")
		}
	})
}

func TestInfoDate_ImportSolarTerm(t *testing.T) {
	orig := solarTermEventTitlesFetcher
	t.Cleanup(func() { solarTermEventTitlesFetcher = orig })

	t.Run("success", func(t *testing.T) {
		solarTermEventTitlesFetcher = func(dt value.DateJp, tempDir string) ([]string, error) {
			return []string{"solar-a", "moon-a"}, nil
		}
		dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
		i := NewInfoDate(dt)
		if err := i.ImportSolarTerm("tempDir"); err != nil {
			t.Fatalf("ImportSolarTerm() error = %v", err)
		}
		if len(i.Events) != 2 || i.Events[0] != "solar-a" || i.Events[1] != "moon-a" {
			t.Fatalf("Events = %v, want [solar-a moon-a]", i.Events)
		}
	})

	t.Run("error", func(t *testing.T) {
		solarTermEventTitlesFetcher = func(dt value.DateJp, tempDir string) ([]string, error) {
			return nil, errors.New("solar fetch error")
		}
		dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
		i := NewInfoDate(dt)
		err := i.ImportSolarTerm("tempDir")
		if err == nil {
			t.Fatal("ImportSolarTerm() error = nil, want non-nil")
		}
		if !strings.Contains(err.Error(), "solar fetch error") {
			t.Fatalf("ImportSolarTerm() error = %v, want contains %q", err, "solar fetch error")
		}
	})
}

func TestInfoDate_ImportCalendar_NilReceiver(t *testing.T) {
	var i *InfoDate
	if err := i.ImportHoliday("tempDir"); !errs.Is(err, ecode.ErrNullPointer) {
		t.Fatalf("nil.ImportHoliday() error = %v, want %v", err, ecode.ErrNullPointer)
	}
	if err := i.ImportSolarTerm("tempDir"); !errs.Is(err, ecode.ErrNullPointer) {
		t.Fatalf("nil.ImportSolarTerm() error = %v, want %v", err, ecode.ErrNullPointer)
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
