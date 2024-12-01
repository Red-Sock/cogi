package cogi

import (
	"os"

	"github.com/Red-Sock/trace-errors"
)

const (
	lineSkip  = "\n"
	separator = " "
)

type Config struct {
	Points []Points
}

type Points struct {
	From string
	To   string
}

func newConfigFromFile(pth string) (*Config, error) {
	cfg := &Config{}

	fl, err := os.Open(pth)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	buff := make([]byte, 512)

	n, err := fl.Read(buff)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	_ = n

	return cfg, nil
}
