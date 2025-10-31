package core

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/vitalykov/bashnya-backend-basic/bashnya-hw4/internal/cli"
)

var (
	errInconsistentFlags = errors.New("only one flag could be set from -c, -d, -u")
	errTooManyArguments  = errors.New("too many arguments")
)

type UniqSeeker struct {
	config cli.Config
}

func onlyOneTrue(values ...bool) bool {
	haveTrue := false
	for _, val := range values {
		if val {
			if haveTrue {
				return false
			}
			haveTrue = true
		}
	}
	return true
}

func NewUniqSeeker(cfg cli.Config) (*UniqSeeker, error) {
	if !onlyOneTrue(cfg.Count, cfg.Repeated, cfg.Unique) {
		return nil, errInconsistentFlags
	}
	seeker := &UniqSeeker{config: cfg}

	return seeker, nil
}

func (us UniqSeeker) SeekUnique(args []string) error {
	l := len(args)
	input := os.Stdin
	output := os.Stdout
	switch {
	case l > 2:
		return errTooManyArguments
	case l == 2:
		outFile, err := os.Open(args[1])
		if err != nil {
			return err
		}
		defer outFile.Close()
		output = outFile
		fallthrough
	case l == 1:
		inFile, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer inFile.Close()
		input = inFile
	}
	sc := bufio.NewScanner(input)
	wr := bufio.NewWriter(output)
	fmt.Println("Scanning")
	for sc.Scan() {
		wr.WriteString(string(sc.Bytes()))
	}
	wr.Flush()
	return nil
}
