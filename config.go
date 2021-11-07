package main

import "flag"

type CatConfig struct {
	filePaths []string
	shouldNumberLines bool
	shouldNumberNonBlankLines bool
	showTabs bool
	squeezeBlank bool
	showEOL bool
	showVersion bool
}

func NewCatConfig() CatConfig {
	config := CatConfig{}
	InitFlags(&config)
	return config
}

func InitFlags(config *CatConfig) {
	flag.BoolVar(&(config.showVersion), "version", false, "output version information and exit")
	flag.BoolVar(&(config.showVersion), "v", false, "output version information and exit")

	flag.Parse()

	config.filePaths = flag.Args()
}