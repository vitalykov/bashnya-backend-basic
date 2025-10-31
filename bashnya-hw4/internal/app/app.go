package app

import (
	"github.com/vitalykov/bashnya-backend-basic/bashnya-hw4/internal/cli"
	"github.com/vitalykov/bashnya-backend-basic/bashnya-hw4/internal/core"
)

const (
	Name = "uniq"
)

type App struct {
	seeker *core.UniqSeeker
}

func NewApp(cfg cli.Config) (*App, error) {
	seeker, err := core.NewUniqSeeker(cfg)
	if err != nil {
		return nil, err
	}
	return &App{seeker: seeker}, nil
}

func Launch(args []string) error {
	cfg, arguments, err := cli.ParseFlags(args)
	if err != nil {
		return err
	}
	app, err := NewApp(cfg)
	if err != nil {
		return nil
	}
	err = app.seeker.SeekUnique(arguments)
	if err != nil {
		return err
	}

	return nil
}
