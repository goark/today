package events

import (
	"encoding/json"
	"io"

	"github.com/goark/errs"
	"github.com/goark/koyomi/value"
	"github.com/goark/today/internal/misc"
)

type Event struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Title string `json:"title"`
}

func ImportEvents(r io.Reader) ([]Event, error) {
	evt := []Event{}
	dec := json.NewDecoder(r)
	for i := 0; dec.More(); i++ {
		var e []Event
		if err := dec.Decode(&e); err != nil {
			return nil, errs.Wrap(err, errs.WithContext("count", i))
		}
		evt = append(evt, e...)
	}
	return evt, nil
}

// Contain checks if the given date is within the event's start and end dates.
func (e Event) Contain(dt value.DateJp) (bool, error) {
	st, serr := misc.DateFrom(e.Start)
	if serr != nil {
		return false, errs.Wrap(serr, errs.WithContext("start", e.Start))
	}

	ed, eerr := misc.DateFrom(e.End)
	if eerr != nil {
		return false, errs.Wrap(eerr, errs.WithContext("end", e.End))
	}
	if dt.Before(st) || dt.After(ed) {
		return false, nil
	}
	return true, nil
}

// GetEventsContaining returns the titles of events that contain the given date.
func GetEventsContaining(dt value.DateJp, events []Event) ([]string, error) {
	result := []string{}
	for _, e := range events {
		contain, err := e.Contain(dt)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		if contain {
			result = append(result, e.Title)
		}
	}
	return result, nil
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
