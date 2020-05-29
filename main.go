package main

import "os"

func main() {

	concaterBuilder := NewConcaterBuilder()
	concaterBuilder.UsingTarget(os.Stdout)
	concaterBuilder.AddSource(os.Stdin)

	concater := concaterBuilder.Build()
	concater.Concatenate()
}
