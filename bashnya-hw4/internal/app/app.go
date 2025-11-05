package app

import (
	"github.com/vitalykov/bashnya-backend-basic/bashnya-hw4/internal/cli"
	"github.com/vitalykov/bashnya-backend-basic/bashnya-hw4/internal/core"
)

const (
	Name = "uniq"
)

// type App struct {
// 	processor *core.UniqProcessor
// }

func LaunchApp(args []string) error {
	parser := cli.NewArgsParser(args)
	cfg, err := parser.GetConfig()
	if err != nil {
		return err
	}
	proc, err := core.NewUniqProcessor(&cfg)
	if err != nil {
		return err
	}
	if err := proc.Process(); err != nil {
		return err
	}
	if err := parser.CloseFiles(); err != nil {
		return err
	}
	return nil
}
