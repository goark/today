package facade

import (
	"fmt"
	"io"
	"time"

	"github.com/goark/today/internal/config"
	"github.com/goark/today/internal/ecode"
	"github.com/goark/today/internal/today"
)

func Execute(in io.Reader, out, errOut io.Writer, version string, args []string) int {
	_ = in
	cfg, err := config.Parse(args, errOut)
	if err != nil {
		if err == config.ErrHelp {
			return int(ecode.Normal)
		}
		fmt.Fprintln(errOut, err)
		return int(ecode.Error)
	}

	if cfg.ShowVersion {
		fmt.Fprintln(out, version)
		return int(ecode.Normal)
	}

	fmt.Fprintln(out, today.String(time.Now()))
	return int(ecode.Normal)
}
