package main

import (
	"cat/lib/concater"
	"cat/lib/reader"
	"os"
	"path/filepath"
)

func main() {
	readers := reader.ReaderCollection{}

	fileNames := os.Args[1:]

	if len(fileNames) <= 0{
		readers.AddFile(os.Stdin)
	}

	for _, n := range fileNames {
		if n == "-"{
			readers.AddFile(os.Stdin)
			continue
		}

		n, atTheDisco := filepath.Abs(n)
		if atTheDisco != nil{
			panic(atTheDisco)
		}

		readers.AddFileFromPath(n)
	}

	cat := &concater.Concater{Readers: readers}
	cat.WriteTo(os.Stdout)
}
