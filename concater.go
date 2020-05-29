package main

import "os"

var (
	BUFFER_SIZE uint16 = 256
	BUFFER      []byte = make([]byte, BUFFER_SIZE)
)

type Concater struct {
	sources []*os.File
	target  *os.File
}

func (c Concater) Concatenate() {
	for {
		nextChunk := c.readNext()
		if len(nextChunk) <= 0 {
			break
		}

		c.write(nextChunk)
	}
}

func (c Concater) readNext() []byte {
	// FixMe multiple sources must be used here
	readedBytes, err := c.sources[0].Read(BUFFER)
	if readedBytes <= 0 {
		return nil
	}

	if err != nil {
		panic(err)
	}

	return BUFFER
}

func (c Concater) write(buffer []byte) {
	_, err := c.target.Write(buffer)
	if err != nil {
		panic(err)
	}
}
