package reader

import (
	"cat/lib/iterator"
	"os"
)

type ReaderCollection struct {
	readers []interface{}
}

func (r *ReaderCollection) AddReader(reader iterator.Iterable)  {
	r.readers = append(r.readers, reader)
}

func (r *ReaderCollection) AddFile(file *os.File)  {
	read := &FileReader{File: file}
	r.AddReader(read)
}

func (r *ReaderCollection) AddFileFromPath(fileName string)  {
	file, e := os.Open(fileName)
	if e != nil{
		panic(e)
	}

	r.AddFile(file)
}

func (r ReaderCollection) GetIterator() iterator.Iterator {
	return iterator.NewIterable(r.readers).GetIterator()
}
