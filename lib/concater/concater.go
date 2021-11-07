package concater

import (
	"cat/lib/reader"
	"os"
)

type Concater struct {
	Readers ReaderCollection
}

func (c Concater) WriteTo(out *os.File) {
	itr := c.Readers.GetIterator()
	moved, _ := itr.Move()
	for moved {
		current := itr.Current().(reader.Reader)
		c.writeTo(current, out)

		moved, _ = itr.Move()
	}
}

func (c Concater) writeTo(reader reader.Reader, out *os.File) {
	itr := reader.GetIterator()
	moved, _ := itr.Move()
	for moved {
		current := itr.Current().([]byte)
		_, e := out.Write(current)
		if e != nil{
			panic(e)
		}

		moved, _ = itr.Move()
	}
}

