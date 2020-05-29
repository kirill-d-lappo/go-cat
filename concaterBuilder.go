package main

import "os"

type ConcaterBuilder struct {
	concater Concater
}

func NewConcaterBuilder() *ConcaterBuilder {
	return &ConcaterBuilder{
		concater: Concater{},
	}
}

func (b *ConcaterBuilder) Build() *Concater {
	return &b.concater
}

func (b *ConcaterBuilder) UsingTarget(target *os.File) *ConcaterBuilder {
	b.concater.target = target
	return b
}

func (b *ConcaterBuilder) AddSource(source *os.File) *ConcaterBuilder {
	b.concater.sources = append(b.concater.sources, source)
	return b
}
