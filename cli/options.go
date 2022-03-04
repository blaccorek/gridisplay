package cli

import (
	"flag"
	"fmt"
)

type FilePathList []string

func (i *FilePathList) String() string {
	s := *i
	return fmt.Sprintln(s)
}

func (i *FilePathList) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type Options struct {
	FilePaths FilePathList
	Portrait  bool
}

func ParseOptions() Options {
	var options = Options{}
	flag.Var(&options.FilePaths, "f", "Image file path (must provide 4 files)")
	options.Portrait = *flag.Bool("portrait", false, "Use portrait orientation images")
	flag.Parse()
	return options
}
