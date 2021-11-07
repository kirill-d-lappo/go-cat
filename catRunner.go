package main

import (
	"cat/lib/concater"
	"cat/lib/reader"
	"os"
	"path/filepath"
)

func RunCat(config CatConfig) {
	readers := reader.ReaderCollection{}

	ApplyConfig(config, readers)

	cat := &concater.Concater{Readers: readers}
	cat.WriteTo(os.Stdout)
}

func ApplyConfig(config CatConfig, readers reader.ReaderCollection) {
	if len(config.filePaths) <= 0 {
		readers.AddFile(os.Stdin)
	}

	paths := config.filePaths
	for _, n := range paths {
		if n == "-" {
			readers.AddFile(os.Stdin)
			continue
		}

		n, atTheDisco := filepath.Abs(n)
		if atTheDisco != nil {
			panic(atTheDisco)
		}

		readers.AddFileFromPath(n)
	}
}