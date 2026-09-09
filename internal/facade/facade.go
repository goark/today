package facade

import (
	"runtime"
	"time"

	"github.com/goark/errs"
	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"github.com/goark/koyomi/value"
	"github.com/goark/struct2pflag"
	"github.com/spf13/pflag"

	"github.com/goark/today/internal/config"
	"github.com/goark/today/internal/misc"
	"github.com/goark/today/internal/today"
)

// Execute is the main function of the application, which executes the command-line interface
func Execute(ui *rwi.RWI, appVersion string, args []string) (exit exitcode.ExitCode) {
	// defer function to catch panic and print stack trace
	defer func() {
		if r := recover(); r != nil {
			if ui != nil {
				_ = ui.OutputErrln("Panic:", r) // ignored error if output fails
				for depth := 0; ; depth++ {
					pc, _, line, ok := runtime.Caller(depth)
					if !ok {
						break
					}
					_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ": line", line)
				}
			}
			exit = exitcode.Abnormal
		}
	}()

	exit = exitcode.Normal
	if err := run(ui, args, NewVersionString(appVersion)); err != nil {
		exit = exitcode.Abnormal
	}
	return
}

func run(ui *rwi.RWI, args []string, ver versionString) error {
	// Parse command-line arguments and bind them to the configuration struct
	fs := pflag.NewFlagSet(usageString(), pflag.ContinueOnError) // Create a new flag set for command-line options
	fs.SetOutput(ui.Writer())
	cfg := config.DefaultConfig(appNameShort)
	struct2pflag.Bind(fs, cfg)             // Bind command-line flags to the configuration struct
	if err := fs.Parse(args); err != nil { // Parse command-line arguments
		if errs.Is(err, pflag.ErrHelp) {
			return nil
		}
		return debugPrint(ui, cfg, errs.Wrap(err))
	}

	// If the version flag is set, print version information and exit
	if cfg.IsVersion() {
		return debugPrint(ui, cfg, errs.Wrap(ui.Outputln(ver)))
	}

	// Determine the date for which to show information. Default is today.
	dt := value.NewDate(time.Now())
	if fs.NArg() > 0 {
		if parsedDate, err := misc.DateFrom(fs.Arg(0)); err == nil {
			dt = parsedDate
		} else {
			return debugPrint(ui, cfg, errs.Wrap(err, errs.WithContext("date", fs.Arg(0))))
		}
	}

	// Show information for the determined date.
	if cfg.IsJSON() {
		if err := today.ShowInformationJSON(ui.Writer(), dt, cfg); err != nil {
			return debugPrint(ui, cfg, errs.Wrap(err, errs.WithContext("date", dt)))
		}
	} else if err := today.ShowInformation(ui.Writer(), dt, cfg); err != nil {
		return debugPrint(ui, cfg, errs.Wrap(err, errs.WithContext("date", dt)))
	}
	return nil
}

func debugPrint(ui *rwi.RWI, cfg *config.Config, err error) error {
	if err == nil {
		return nil
	}
	if cfg.IsDebug() {
		if perr := ui.Outputln(errs.EncodeJSON(err)); perr != nil {
			err = errs.Join(err, perr)
		}
		return err
	}
	if perr := ui.OutputErrln("Error:", err); perr != nil {
		err = errs.Join(err, perr)
	}
	return err // Return the error after attempting to print it.
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
