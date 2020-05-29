package concater

import "os"

type ConcaterBuilder struct {
	concater Concater
}

func NewConcaterBuilder() *ConcaterBuilder {
	return &ConcaterBuilder{
		concater: Concater{
			nextSourceIndex: 0,
		},
	}
}

func (b *ConcaterBuilder) Build() *Concater {
	return &b.concater
}

func (b *ConcaterBuilder) UsingTarget(target *os.File) *ConcaterBuilder {
	b.concater.target = target
	return b
}

func (b *ConcaterBuilder) AddSource(source func() *os.File) *ConcaterBuilder {
	b.concater.openSourceActions = append(b.concater.openSourceActions, source)
	return b
}
