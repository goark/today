package facade

import (
	"fmt"
	"runtime/debug"
	"strings"
)

const (
	appNameShort  = "today"                            // Name of the application
	appNameLong   = "github.com/goark/" + appNameShort // Long name of the application
	repositoryURL = "https://" + appNameLong           // Repository URL of the application
)

// versionString is a type for version information of the application
type versionString []string

// NewVersionString returns a new versionString instance with the given version string
func NewVersionString(v string) versionString {
	return versionString([]string{
		strings.Join([]string{appNameShort, replaceVersion(v)}, " "),
		repositoryURL,
	})
}

// replaceVersion returns the version string based on the provided version and build information
func replaceVersion(v string) string {
	if v != "" {
		return v
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		goVersion := fmt.Sprintf("(compiled with %v)", info.GoVersion)
		if info.Main.Version != "" {
			return joinNonEmpty(info.Main.Version, goVersion)
		}
		var revision, dirty string
		for _, v := range info.Settings {
			switch v.Key {
			case "vcs.revision":
				revision = v.Value
			case "vcs.modified":
				if v.Value == "true" {
					dirty = "(dirty)"
				}
			}
		}
		if revision != "" {
			return joinNonEmpty(revision, dirty, goVersion)
		}
	}
	return "(version not set)"
}

func joinNonEmpty(parts ...string) string {
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p != "" {
			out = append(out, p)
		}
	}
	return strings.Join(out, " ")
}

// String returns the string representation of the versionString
func (v versionString) String() string {
	return strings.Join(v, "\n")
}

// UsageString returns the usage string of header part"
func usageString() string {
	return strings.Join([]string{
		appNameShort,
		"",
		"Usage:",
		"  " + appNameShort + " [flags] [yyyy-mm-dd]",
		"",
		"Flags",
	}, "\n")
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
