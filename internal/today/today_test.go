package today

import (
	"bytes"
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
