package config

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/goark/errs"
	"github.com/goark/today/internal/ecode"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig("today")
	if cfg == nil {
		t.Fatal("DefaultConfig() = nil, want non-nil")
	}
	if cfg.VersionFlag {
		t.Fatal("VersionFlag = true, want false")
	}
	if cfg.DebugFlag {
		t.Fatal("DebugFlag = true, want false")
	}
	if cfg.EventFile == "" {
		t.Fatal("EventFile is empty")
	}
	if !strings.HasSuffix(cfg.EventFile, "/today/event.json") {
		t.Fatalf("EventFile = %q, want suffix %q", cfg.EventFile, "/today/event.json")
	}
}

func TestConfigFlagHelpers(t *testing.T) {
	var nilCfg *Config
	if nilCfg.IsVersion() || nilCfg.IsDebug() || nilCfg.IsHoliday() || nilCfg.IsSolarTerm() || nilCfg.IsOtherEvents() {
		t.Fatal("nil config helper should return false")
	}

	cfg := &Config{}
	if cfg.IsVersion() || cfg.IsDebug() || cfg.IsHoliday() || cfg.IsSolarTerm() || cfg.IsOtherEvents() {
		t.Fatal("default helper flags should be false")
	}

	cfg.VersionFlag = true
	cfg.DebugFlag = true
	cfg.HolidayFlag = true
	cfg.SolarTermFlag = true
	cfg.OtherEventsFlag = true
	if !cfg.IsVersion() || !cfg.IsDebug() || !cfg.IsHoliday() || !cfg.IsSolarTerm() || !cfg.IsOtherEvents() {
		t.Fatal("helper flags should reflect explicit flags")
	}

	cfg = &Config{AllEventsFlag: true}
	if !cfg.IsHoliday() || !cfg.IsSolarTerm() || !cfg.IsOtherEvents() {
		t.Fatal("AllEventsFlag should enable holiday/solar-term/other-events")
	}
}

func TestOpenEvent(t *testing.T) {
	var nilCfg *Config
	if _, err := nilCfg.OpenEvent(); !errs.Is(err, ecode.ErrNoEventFile) {
		t.Fatalf("nilCfg.OpenEvent() error = %v, want %v", err, ecode.ErrNoEventFile)
	}

	emptyPath := &Config{}
	if _, err := emptyPath.OpenEvent(); !errs.Is(err, ecode.ErrNoEventFile) {
		t.Fatalf("emptyPath.OpenEvent() error = %v, want %v", err, ecode.ErrNoEventFile)
	}

	missing := &Config{EventFile: t.TempDir() + "/missing.json"}
	if _, err := missing.OpenEvent(); !errs.Is(err, ecode.ErrNoEventFile) {
		t.Fatalf("missing.OpenEvent() error = %v, want %v", err, ecode.ErrNoEventFile)
	}

	tmp := t.TempDir() + "/event.json"
	if err := os.WriteFile(tmp, []byte("[]"), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	good := &Config{EventFile: tmp}
	r, err := good.OpenEvent()
	if err != nil {
		t.Fatalf("good.OpenEvent() error = %v", err)
	}
	defer r.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("ReadAll() error = %v", err)
	}
	if string(b) != "[]" {
		t.Fatalf("event file content = %q, want %q", string(b), "[]")
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
