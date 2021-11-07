package reader

import (
	"cat/lib/iterator"
	"os"
)

var (
	BufferSize int    = 256
	Buffer     []byte = make([]byte, BufferSize)
)

type FileReader struct {
	File *os.File
	current []byte
}

func (fr FileReader) Move() (bool, error) {
	nextBytes, e := fr.readNextBytes()
	if e != nil {
		fr.dispose()
		return false, e
	}

	if len(nextBytes) <= 0{
		fr.dispose()
		return false, nil
	}

	fr.current = nextBytes
	return true, nil
}

func (fr FileReader) Current() interface{} {
	return fr.current
}

func (fr *FileReader) GetIterator() iterator.Iterator {
	return fr
}

func (fr *FileReader) readNextBytes() ([]byte, error) {

	readBytesAmount, err := fr.File.Read(Buffer)
	if err != nil {
		return nil, err
	}

	return Buffer[:readBytesAmount], nil
}

func (fr *FileReader) dispose()  {
	if fr.File != nil {
		_ = fr.File.Close()
	}
}