package main

import (
	"github.com/jessevdk/go-flags"
)

type CatOptions struct {
	ShowNonBlankLineNumber bool `short:"b" long:"number-nonblank" description:"number nonempty output lines, overrides -n"`
	ShowEOL                bool `short:"E" long:"show-ends" description:"display $ at end of each line"`
	ShowLineNumber         bool `short:"n" long:"number" description:"number all output lines"`
	SqueezeBlank           bool `short:"s" long:"squeeze-blank" description:"suppress repeated empty output lines"`
	ShowTabs               bool `short:"T" long:"show-tabs" description:"display TAB characters as ^I"`
	ShowNonPrinting        bool `short:"v" long:"show-nonprinting" description:"use ^ and M- notation, except for LFD and TAB"`
	ShowVersion            bool `long:"version" description:"output version information and exit"`
	ShowAll                bool `short:"A" long:"show-all" description:"equivalent to -vET"`
	ShowNonPrintingEOL     bool `short:"e" long:"" description:"equivalent to -vE"`
	IgnoredU               bool `short:"u" description:"(ignored)"`
}

type CatConfig struct {
	filePaths []string
	options   CatOptions
}

func NewCatConfig(args []string) (*CatConfig, error) {
	options := CatOptions{}
	parsedArgs, e := flags.ParseArgs(&options, args)
	if e != nil {
		return nil, e
	}

	config := &CatConfig{
		filePaths: parsedArgs,
		options:   options,
	}

	return config, nil
}
