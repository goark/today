package config

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

var ErrHelp = errors.New("help requested")

type Config struct {
	ShowVersion bool
}

func Parse(args []string, errOut io.Writer) (*Config, error) {
	fs := flag.NewFlagSet("today", flag.ContinueOnError)
	fs.SetOutput(errOut)

	cfg := &Config{}
	fs.BoolVar(&cfg.ShowVersion, "version", false, "show version")

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil, fmt.Errorf("%w", ErrHelp)
		}
		return nil, err
	}

	return cfg, nil
}
