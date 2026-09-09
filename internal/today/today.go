package today

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/goark/errs"
	"github.com/goark/koyomi/value"
	"github.com/goark/today/internal/config"
	"github.com/goark/today/internal/ecode"
)

// ShowInformationJSON writes the detailed information of the given date in JSON format to the provided writer.
// It returns an error if the date is zero or if writing to the writer fails.
func ShowInformationJSON(w io.Writer, dt value.DateJp, cfg *config.Config) error {
	// Get the detailed information of the date in JSON format.
	i, err := informationJSON(w, dt, cfg)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("date", dt))
	}

	// write the detailed information of the date in JSON format to the provided writer
	if err := json.NewEncoder(w).Encode(i); err != nil {
		return errs.Wrap(err, errs.WithContext("date", dt), errs.WithContext("info", i))
	}
	return nil
}

// ShowInformation writes the detailed information of the given date to the provided writer.
// It returns an error if the date is zero or if writing to the writer fails.
func ShowInformation(w io.Writer, dt value.DateJp, cfg *config.Config) error {
	// Get the detailed information of the date in JSON format.
	i, err := informationJSON(w, dt, cfg)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("date", dt))
	}
	// write the detailed information of the date to the provided writer
	if _, err := fmt.Fprintln(w, i); err != nil {
		return errs.Wrap(err, errs.WithContext("date", dt), errs.WithContext("info", i))
	}
	return nil
}

// informationJSON retrieves the detailed information of the given date in JSON format.
// It returns an error if the date is zero or if importing any information fails.
func informationJSON(w io.Writer, dt value.DateJp, cfg *config.Config) (info *InfoDate, err error) {
	// Check if the provided date is zero. If it is, return an error.
	if dt.IsZero() {
		return nil, errs.Wrap(ecode.ErrZeroValue)
	}
	// Create a new InfoDate instance for the given date.
	i := NewInfoDate(dt)

	// import holiday information if the holiday flag is set
	if cfg.IsHoliday() {
		if err := i.ImportHoliday(cfg.TempDir); err != nil {
			return nil, errs.Wrap(err, errs.WithContext("date", dt))
		}
	}

	// import solar term information if the solar term flag is set
	if cfg.IsSolarTerm() { // import solar term information if the solar term flag is set
		if err := i.ImportSolarTerm(cfg.TempDir); err != nil {
			return nil, errs.Wrap(err, errs.WithContext("date", dt))
		}
	}

	// import other events if the other events flag is set
	if cfg.IsOtherEvents() { // import events if the events flag is set
		r, oerr := cfg.OpenEvent()
		if oerr != nil && !errs.Is(oerr, ecode.ErrNoEventFile) {
			return nil, errs.Wrap(oerr, errs.WithContext("date", dt), errs.WithContext("file", cfg.EventFile))
		}
		if r != nil {
			defer func() {
				if r != nil {
					cerr := r.Close()
					if cerr != nil {
						err = errs.Join(err, errs.Wrap(cerr, errs.WithContext("date", dt), errs.WithContext("file", cfg.EventFile)))
					}
				}
			}()
			if err := i.ImportEvents(r); err != nil {
				return nil, errs.Wrap(err, errs.WithContext("date", dt), errs.WithContext("file", cfg.EventFile))
			}
		}
	}

	return i, nil
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
