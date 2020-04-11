package flag

import (
	"errors"
	"flag"
	"path/filepath"
)

type Config struct {
	InputGoFilePath  string
	OutputGoFilePath string
	Verbose          bool
}

func ShellParsing( /*os.Args as is*/ args []string) (conf Config, _ error) {
	usage := "TODO: как использовать эту улилиту" // TODO
	flag.CommandLine = flag.NewFlagSet(usage, flag.ContinueOnError)

	outputGoFilePath := flag.CommandLine.String("output_filename", "\"${FILENAME}_dao.go\"", "specify the filename of the output instead")
	verbose := flag.CommandLine.Bool("verbose", false, "detailed process logging")

	if err := flag.CommandLine.Parse(args[1:]); err != nil {
		return Config{}, err
	}

	conf.OutputGoFilePath = *outputGoFilePath
	conf.Verbose = *verbose

	if args := flag.CommandLine.Args(); len(args) == 0 {
		return Config{}, errors.New("insufficient arguments: the target file is not specified")
	} else if len(args) > 1 {
		return Config{}, errors.New("too many arguments - specify 1 target file")
	} else {
		conf.InputGoFilePath = args[0]
	}

	extention := filepath.Ext(conf.InputGoFilePath)
	if extention == ".go" {
		conf.OutputGoFilePath = conf.InputGoFilePath[0:len(conf.InputGoFilePath)-len(".go")] + "_dao.go"
	} else {
		return Config{}, errors.New("not valid \"*.go\" file: " + conf.InputGoFilePath)
	}

	return conf, nil
}
