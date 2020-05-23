package main

import (
	"os"
)

func main() {
	sourceFile := os.Stdin

	outputFile := os.Stdout

	buffer := make([]byte, 256)

	for {
		readBytes, err := sourceFile.Read(buffer)
		if readBytes <= 0 {
			return
		}

		if err != nil {
			panic(err)
		}

		_, err = outputFile.Write(buffer)
		if err != nil {
			panic(err)
		}
	}
}
