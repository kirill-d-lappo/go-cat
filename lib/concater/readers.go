package concater

import (
	"cat/lib/iterator"
	"cat/lib/reader"
	"os"
)

type ReaderCollection struct {
	readers []iterator.Iterable
	currentReaderIndex int
}

func (r *ReaderCollection) AddReader(reader iterator.Iterable)  {
	r.readers = append(r.readers, reader)
}

func (r *ReaderCollection) AddFile(file *os.File)  {
	read := &reader.FileReader{File: file}
	r.AddReader(read)
}

func (r *ReaderCollection) AddFileFromPath(fileName string)  {
	file, e := os.Open(fileName)
	if e != nil{
		panic(e)
	}

	r.AddFile(file)
}

func (r ReaderCollection) Move() (bool, error) {
	if r.currentReaderIndex >= len(r.readers){
		return false, nil
	}

	r.currentReaderIndex++
	return true, nil
}

func (r ReaderCollection) Current() interface{} {
	return r.readers[r.currentReaderIndex]
}

func (r ReaderCollection) GetIterator() iterator.Iterator {
	return r
}
