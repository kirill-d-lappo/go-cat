package main

import (
	"cat/lib/concater"
	"cat/lib/reader"
	"os"
	"path/filepath"
)

func RunCat(config *CatConfig) {
	if config.options.ShowVersion {
		printVersion()
		return
	}

	readers := reader.ReaderCollection{}

	ApplyConfig(config, &readers)

	cat := &concater.Concater{Readers: readers}
	cat.WriteTo(os.Stdout)
}

func ApplyConfig(config *CatConfig, readers *reader.ReaderCollection) {
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

func printVersion() {
	versionText := `cat (GO-CAT project) 0.1.0
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Original cat is written by Torbjorn Granlund and Richard M. Stallman.

go version is written by Kirill Lappo as a part of go-cat project.

just for fun
`

	print(versionText)
}
