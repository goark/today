package misc

import (
	"strings"
	"time"

	"github.com/goark/errs"
	"github.com/goark/koyomi/value"
)

// DateFrom parses a string into a value.DateJp using predefined time templates.
// If the string is empty or "null", it returns a zero value date.
// It returns an error if the string cannot be parsed with any of the templates.
func DateFrom(s string) (value.DateJp, error) {
	zero := value.NewDate(time.Time{})
	if len(s) == 0 || strings.EqualFold(s, "null") {
		return zero, nil
	}
	errlist := &errs.Errors{}
	for _, tmplt := range timeTemplate {
		if tm, err := time.Parse(tmplt, s); err != nil {
			errlist.Add(errs.Wrap(err, errs.WithContext("time_string", s), errs.WithContext("time_template", tmplt)))
		} else {
			return value.NewDate(tm), nil
		}
	}
	return zero, errlist.ErrorOrNil()
}

// timeTemplate defines the list of time formats used for parsing date strings.
var timeTemplate = []string{
	"2006-01-02",
	"20060102",
	time.RFC3339,
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
