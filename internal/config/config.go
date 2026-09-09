package config

import (
	"io"
	"os"

	"github.com/goark/errs"
	"github.com/goark/gocli/cache"
	cfg "github.com/goark/gocli/config"
	"github.com/goark/today/internal/ecode"
)

// Config represents toptags command options.
type Config struct {
	VersionFlag     bool   `pflag:"version,v,show version information"`
	DebugFlag       bool   `pflag:"debug,,enable debug mode"`
	HolidayFlag     bool   `pflag:"holiday,h,show holiday information"`
	SolarTermFlag   bool   `pflag:"solar-term,s,show solar term information"`
	OtherEventsFlag bool   `pflag:"other-events,e,show other events information"`
	AllEventsFlag   bool   `pflag:"all-events,a,show all events information"`
	EventFile       string `pflag:"event-file,,path to the event file"`
	JSONFlag        bool   `pflag:"json,j,output information in JSON format"`
	TempDir         string `pflag:"temp-dir,,path to the temporary directory"`
}

// DefaultConfig returns a default Config instance.
func DefaultConfig(appName string) *Config {
	return &Config{
		EventFile: cfg.Path(appName, "event.json"),
		TempDir:   cache.Dir(appName),
	}
}

// IsVersion returns true if the version flag is set.
func (c *Config) IsVersion() bool {
	if c == nil {
		return false
	}
	return c.VersionFlag
}

// IsDebug returns true if the debug flag is set.
func (c *Config) IsDebug() bool {
	if c == nil {
		return false
	}
	return c.DebugFlag
}

// IsHoliday returns true if the holiday flag is set.
func (c *Config) IsHoliday() bool {
	if c == nil {
		return false
	}
	return c.AllEventsFlag || c.HolidayFlag
}

// IsJSON returns true if the JSON flag is set.
func (c *Config) IsJSON() bool {
	if c == nil {
		return false
	}
	return c.JSONFlag
}

// IsSolarTerm returns true if the solar term flag is set.
func (c *Config) IsSolarTerm() bool {
	if c == nil {
		return false
	}
	return c.AllEventsFlag || c.SolarTermFlag
}

// IsOtherEvents returns true if the other events flag is set.
func (c *Config) IsOtherEvents() bool {
	if c == nil {
		return false
	}
	return c.AllEventsFlag || c.OtherEventsFlag
}

// OpenEvent opens the event file and returns a ReadCloser. If the event file is not specified or cannot be opened, ErrNoEventFile error is returned.
func (c *Config) OpenEvent() (io.ReadCloser, error) {
	if c == nil || c.EventFile == "" {
		return nil, errs.Wrap(ecode.ErrNoEventFile, errs.WithContext("file", ""))
	}
	f, err := os.Open(c.EventFile)
	if err != nil {
		if errs.Is(err, os.ErrNotExist) {
			err = errs.Wrap(ecode.ErrNoEventFile)
		}
		return nil, errs.Wrap(err, errs.WithContext("file", c.EventFile))
	}
	return f, nil
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
