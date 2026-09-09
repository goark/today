package facade

import (
	"bytes"
	"strings"
	"testing"

	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
)

func TestExecute_NilUI(t *testing.T) {
	got := Execute(nil, "v1.2.3", []string{"--version"})
	if got != exitcode.Abnormal {
		t.Fatalf("Execute(nil, ...) = %v, want %v", got, exitcode.Abnormal)
	}
}

func TestExecute_JSONOutput(t *testing.T) {
	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	ui := rwi.New(
		rwi.WithWriter(out),
		rwi.WithErrorWriter(errOut),
	)

	got := Execute(ui, "v1.2.3", []string{"--json", "2026-09-09"})
	if got != exitcode.Normal {
		t.Fatalf("Execute(..., --json) = %v, want %v, stderr=%q", got, exitcode.Normal, errOut.String())
	}
	if !strings.Contains(out.String(), "\"year\":2026") {
		t.Fatalf("stdout = %q, want contains %q", out.String(), "\"year\":2026")
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
