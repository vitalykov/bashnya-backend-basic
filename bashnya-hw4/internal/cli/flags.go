package cli

import (
	"flag"
)

type Config struct {
	Count      bool
	Repeated   bool
	Unique     bool
	IgnoreCase bool
	SkipFields int
	SkipChars  int
}

func ParseFlags(args []string) (Config, []string, error) {
	var cfg Config
	flags := flag.NewFlagSet("uniq", flag.ExitOnError)

	flags.BoolVar(&cfg.Count, "c", false, "prefix lines by the number of occurrences")
	flags.BoolVar(&cfg.Repeated, "d", false, "only print duplicate lines, one for each group")
	flags.BoolVar(&cfg.Unique, "u", false, "only print unique lines")
	flags.BoolVar(&cfg.IgnoreCase, "i", false, "ignore differences in case when comparing")
	flags.IntVar(&cfg.SkipFields, "f", 0, "avoid comparing the first N fields")
	flags.IntVar(&cfg.SkipChars, "s", 0, "avoid comparing the first N characters")

	if err := flags.Parse(args[1:]); err != nil {
		return Config{}, []string{}, err
	}

	nonFlags := flags.Args()
	return cfg, nonFlags, nil
}
