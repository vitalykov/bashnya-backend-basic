package cli

import (
	"errors"
	"flag"
	"os"

	"github.com/vitalykov/bashnya-backend-basic/bashnya-hw4/internal/core"
)

var errTooManyArguments = errors.New("too many arguments")

type ArgsParser struct {
	args    []string
	inFile  *os.File
	outFile *os.File
}

func NewArgsParser(args []string) *ArgsParser {
	return &ArgsParser{args, os.Stdin, os.Stdout}
}

func (parser ArgsParser) GetConfig() (core.Config, error) {
	var opts core.Options
	flags := flag.NewFlagSet("uniq", flag.ExitOnError)

	flags.BoolVar(&opts.Count, "c", false, "prefix lines by the number of occurrences")
	flags.BoolVar(&opts.Repeated, "d", false, "only print duplicate lines, one for each group")
	flags.BoolVar(&opts.Unique, "u", false, "only print unique lines")
	flags.BoolVar(&opts.IgnoreCase, "i", false, "ignore differences in case when comparing")
	flags.IntVar(&opts.SkipFields, "f", 0, "avoid comparing the first N fields")
	flags.IntVar(&opts.SkipChars, "s", 0, "avoid comparing the first N characters")

	if err := flags.Parse(parser.args[1:]); err != nil {
		return core.Config{}, err
	}

	nonFlags := flags.Args()
	params, err := parser.getInputOutput(nonFlags)
	if err != nil {
		return core.Config{}, err
	}

	return core.Config{Options: opts, IOParams: *params}, nil
}

func (parser *ArgsParser) CloseFiles() error {
	if err := parser.inFile.Close(); err != nil {
		return err
	}
	if err := parser.outFile.Close(); err != nil {
		return err
	}
	return nil
}

func (parser *ArgsParser) getInputOutput(args []string) (*core.IOParams, error) {
	l := len(args)
	params := core.IOParams{Input: os.Stdin, Output: os.Stdout}
	switch {
	case l > 2:
		return &params, errTooManyArguments
	case l == 2:
		outFile, err := os.Create(args[1])
		if err != nil {
			return &params, err
		}
		params.Output = outFile
		parser.outFile = outFile
		fallthrough
	case l == 1:
		inFile, err := os.Open(args[0])
		if err != nil {
			return &params, err
		}
		params.Input = inFile
		parser.inFile = inFile
	}
	return &params, nil
}
