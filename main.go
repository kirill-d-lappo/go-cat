package main

import (
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	config, e := NewCatConfig(os.Args[1:])
	if e == nil {
		RunCat(config)
		return
	}

	handleError(e)
}

func handleError(e error) {
	flagError, ok := e.(*flags.Error)
	if ok && flagError.Type == flags.ErrHelp {
		return
	}

	panic(e)
}
