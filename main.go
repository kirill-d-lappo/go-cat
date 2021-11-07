package main

import (
	"cat/lib/concater"
	"os"
)

func main() {
	readers := concater.ReaderCollection{}

	fileNames := os.Args[1:]

	if len(fileNames) <= 0{
		readers.AddFile(os.Stdin)
	}

	for _, n := range fileNames {
		if n == "-"{
			readers.AddFile(os.Stdin)
			continue
		}

		readers.AddFileFromPath(n)
	}

	cat := &concater.Concater{Readers: readers}
	cat.WriteTo(os.Stdout)
}
