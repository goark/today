package today

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/goark/errs"
	"github.com/goark/koyomi/value"
	"github.com/goark/today/internal/config"
	"github.com/goark/today/internal/ecode"
)

func TestShowInformation_ZeroDate(t *testing.T) {
	err := ShowInformation(&bytes.Buffer{}, value.NewDate(time.Time{}), &config.Config{})
	if err == nil {
		t.Fatal("ShowInformation() error = nil, want non-nil")
	}
	if !errs.Is(err, ecode.ErrZeroValue) {
		t.Fatalf("ShowInformation() error = %v, want %v", err, ecode.ErrZeroValue)
	}
}

func TestShowInformation_BasicOutput(t *testing.T) {
	dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
	buf := &bytes.Buffer{}
	err := ShowInformation(buf, dt, &config.Config{})
	if err != nil {
		t.Fatalf("ShowInformation() error = %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "2026-09-09") {
		t.Fatalf("output = %q, want contains %q", out, "2026-09-09")
	}
}

func TestShowInformation_OtherEventsFromFile(t *testing.T) {
	dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
	tmp := t.TempDir() + "/event.json"
	content := `[{"start":"2026-09-01","end":"2026-09-10","title":"event-a"}]`
	if err := os.WriteFile(tmp, []byte(content), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	cfg := &config.Config{OtherEventsFlag: true, EventFile: tmp}
	buf := &bytes.Buffer{}
	err := ShowInformation(buf, dt, cfg)
	if err != nil {
		t.Fatalf("ShowInformation() error = %v", err)
	}
	if !strings.Contains(buf.String(), "event-a") {
		t.Fatalf("output = %q, want contains %q", buf.String(), "event-a")
	}
}

func TestShowInformationJSON_BasicOutput(t *testing.T) {
	dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
	buf := &bytes.Buffer{}
	err := ShowInformationJSON(buf, dt, &config.Config{})
	if err != nil {
		t.Fatalf("ShowInformationJSON() error = %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatalf("JSON unmarshal error = %v, output=%q", err, buf.String())
	}
	if got["year"] != float64(2026) {
		t.Fatalf("year = %v, want 2026", got["year"])
	}
	if got["month"] != float64(9) {
		t.Fatalf("month = %v, want 9", got["month"])
	}
	if got["day"] != float64(9) {
		t.Fatalf("day = %v, want 9", got["day"])
	}
}

func TestShowInformationJSON_TempDirPropagation(t *testing.T) {
	origHoliday := holidayEventTitlesFetcher
	origSolar := solarTermEventTitlesFetcher
	t.Cleanup(func() {
		holidayEventTitlesFetcher = origHoliday
		solarTermEventTitlesFetcher = origSolar
	})

	gotHolidayTempDir := ""
	gotSolarTempDir := ""
	holidayEventTitlesFetcher = func(dt value.DateJp, tempDir string) ([]string, error) {
		gotHolidayTempDir = tempDir
		return []string{"holiday-a"}, nil
	}
	solarTermEventTitlesFetcher = func(dt value.DateJp, tempDir string) ([]string, error) {
		gotSolarTempDir = tempDir
		return []string{"solar-a"}, nil
	}

	dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
	cfg := &config.Config{HolidayFlag: true, SolarTermFlag: true, TempDir: "/tmp/today-test"}
	buf := &bytes.Buffer{}
	err := ShowInformationJSON(buf, dt, cfg)
	if err != nil {
		t.Fatalf("ShowInformationJSON() error = %v", err)
	}
	if gotHolidayTempDir != "/tmp/today-test" {
		t.Fatalf("holiday tempDir = %q, want %q", gotHolidayTempDir, "/tmp/today-test")
	}
	if gotSolarTempDir != "/tmp/today-test" {
		t.Fatalf("solar tempDir = %q, want %q", gotSolarTempDir, "/tmp/today-test")
	}
	if !strings.Contains(buf.String(), "holiday-a") || !strings.Contains(buf.String(), "solar-a") {
		t.Fatalf("output = %q, want contains holiday-a and solar-a", buf.String())
	}
}

func TestShowInformationJSON_HolidayError(t *testing.T) {
	origHoliday := holidayEventTitlesFetcher
	t.Cleanup(func() { holidayEventTitlesFetcher = origHoliday })
	holidayEventTitlesFetcher = func(dt value.DateJp, tempDir string) ([]string, error) {
		return nil, errors.New("holiday fetch error")
	}

	dt := value.NewDate(time.Date(2026, time.September, 9, 0, 0, 0, 0, time.Local))
	cfg := &config.Config{HolidayFlag: true, TempDir: "/tmp/today-test"}
	err := ShowInformationJSON(&bytes.Buffer{}, dt, cfg)
	if err == nil {
		t.Fatal("ShowInformationJSON() error = nil, want non-nil")
	}
	if !strings.Contains(err.Error(), "holiday fetch error") {
		t.Fatalf("ShowInformationJSON() error = %v, want contains %q", err, "holiday fetch error")
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
