package core

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	errInconsistentFlags = errors.New("only one flag could be set from -c, -d, -u")
)

type IOParams struct {
	Input  io.Reader
	Output io.Writer
}

type Options struct {
	Count      bool
	Repeated   bool
	Unique     bool
	IgnoreCase bool
	SkipFields int
	SkipChars  int
}

type Config struct {
	Options
	IOParams
}

type UniqProcessor struct {
	config *Config
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

func NewUniqProcessor(cfg *Config) (*UniqProcessor, error) {
	if !onlyOneTrue(cfg.Count, cfg.Repeated, cfg.Unique) {
		return nil, errInconsistentFlags
	}
	proc := &UniqProcessor{config: cfg}

	return proc, nil
}

func (up UniqProcessor) Process() error {
	sc := bufio.NewScanner(up.config.Input)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return err
	}
	result := SeekUnique(lines, &up.config.Options)
	wr := bufio.NewWriter(up.config.Output)
	for _, l := range result {
		_, err := wr.WriteString(l + "\n")
		if err != nil {
			return err
		}
	}
	if err := wr.Flush(); err != nil {
		return err
	}
	return nil
}

func SeekUnique(lines []string, opts *Options) []string {
	if len(lines) == 0 {
		return []string{}
	}
	var lineConv func(string) string
	if opts.IgnoreCase {
		lineConv = strings.ToLower
	} else {
		lineConv = func(s string) string { return s }
	}
	var lineCount func(string, int) string
	if opts.Count {
		lineCount = func(s string, cnt int) string {
			return strconv.Itoa(cnt) + " " + s
		}
	} else {
		lineCount = func(s string, _ int) string { return s }
	}
	count := 1
	result := []string{}
	toSave := lines[0]
	prevLine := skipped(lineConv(lines[0]), opts.SkipFields, opts.SkipChars)
	for i := 1; i < len(lines); i++ {
		l := skipped(lineConv(lines[i]), opts.SkipFields, opts.SkipChars)
		if l != prevLine {
			if (opts.Repeated && count > 1) || (opts.Unique && count == 1) || (!opts.Unique && !opts.Repeated) {
				result = append(result, lineCount(toSave, count))
			}
			toSave = lines[i]
			prevLine = l
			count = 1
		} else {
			count++
		}
	}
	if (opts.Repeated && count > 1) || (opts.Unique && count == 1) || (!opts.Unique && !opts.Repeated) {
		result = append(result, lineCount(toSave, count))
	}
	return result
}

func skipped(s string, fields int, chars int) string {
	for range fields {
		letterPos := strings.IndexFunc(s, func(r rune) bool {
			return r != ' '
		})
		if letterPos == -1 {
			return ""
		}
		s = s[letterPos:]
		spacePos := strings.IndexByte(s, ' ')
		if spacePos == -1 {
			return ""
		}
		s = s[spacePos:]
	}
	if len(s) < chars {
		return ""
	}
	return s[chars:]
}
