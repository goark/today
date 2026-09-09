package misc

import (
	"testing"
	"time"
)

func TestDateFrom(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		isZero  bool
	}{
		{name: "yyyy-mm-dd", input: "2026-09-09", wantErr: false, isZero: false},
		{name: "yyyymmdd", input: "20260909", wantErr: false, isZero: false},
		{name: "rfc3339", input: "2026-09-09T12:34:56+09:00", wantErr: false, isZero: false},
		{name: "empty", input: "", wantErr: false, isZero: true},
		{name: "null", input: "null", wantErr: false, isZero: true},
		{name: "invalid", input: "2026/09/09", wantErr: true, isZero: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := DateFrom(tc.input)
			if (err != nil) != tc.wantErr {
				t.Fatalf("DateFrom(%q) error = %v, wantErr %v", tc.input, err, tc.wantErr)
			}
			if got.IsZero() != tc.isZero {
				t.Fatalf("DateFrom(%q).IsZero() = %v, want %v", tc.input, got.IsZero(), tc.isZero)
			}
		})
	}
}

func TestDateFrom_ParsedDateValue(t *testing.T) {
	got, err := DateFrom("2026-09-09")
	if err != nil {
		t.Fatalf("DateFrom() error = %v", err)
	}
	if got.Year() != 2026 || got.Month() != time.September || got.Day() != 9 {
		t.Fatalf("DateFrom() = %04d-%02d-%02d, want 2026-09-09", got.Year(), got.Month(), got.Day())
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
