package main

import (
	"cat/lib/concater"
	"os"
)

func main() {

	concaterBuilder := concater.NewConcaterBuilder()
	concaterBuilder.UsingTarget(os.Stdout)

	// FixMe [2020/05/29 KL] Implement argument consumption
	// concaterBuilder.AddSource(func() *os.File { return os.Stdin })
	concaterBuilder.AddSource(func() *os.File { f, _ := os.Open("in1"); return f })
	concaterBuilder.AddSource(func() *os.File { f, _ := os.Open("in2"); return f })

	concater := concaterBuilder.Build()
	concater.Concatenate()
}
