package config

import (
	cfg "github.com/goark/gocli/config"
)

// Config represents toptags command options.
type Config struct {
	VersionFlag bool   `pflag:"version,v,show version information"`
	DebugFlag   bool   `pflag:"debug,,enable debug mode"`
	EventFile   string `pflag:"event-file,,path to the event file"`
}

// DefaultConfig returns a default Config instance.
func DefaultConfig(appName string) *Config {
	return &Config{
		VersionFlag: false,
		DebugFlag:   false,
		EventFile:   cfg.Path(appName, "event.json"),
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
