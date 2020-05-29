package concater

import (
	"os"
)

var (
	BUFFER_SIZE int    = 256
	BUFFER      []byte = make([]byte, BUFFER_SIZE)
)

type Concater struct {
	openSourceActions []func() *os.File
	nextSourceIndex   int
	currentSource     *os.File
	target            *os.File
}

// FixMe [2020/05/29 KL] Review the code, it's overcomplicated. Another approach must be found
func (c *Concater) Concatenate() {
	for c.hasSourcesToRead() {
		nextChunk := c.readNext()
		if len(nextChunk) <= 0 {
			break
		}

		c.write(nextChunk)
	}
}

func (c *Concater) hasSourcesToRead() bool {
	return c.nextSourceIndex < len(c.openSourceActions)
}

func (c *Concater) readNext() []byte {

	buffer, err := c.readNextBytes()
	bufferLen := len(buffer)

	if bufferLen == 0 {
		c.closeCurrentSource()
		buffer, err = c.readNextBytes()
		bufferLen = len(buffer)
	}

	if err != nil {
		panic(err)
	}

	if bufferLen < BUFFER_SIZE {
		c.closeCurrentSource()
	}

	return buffer
}

func (c *Concater) readNextBytes() ([]byte, error) {
	file := c.getCurrentSource()
	if file == nil {
		return nil, nil
	}

	readedBytesAmount, err := file.Read(BUFFER)
	if err != nil {
		return nil, err
	}

	return BUFFER[:readedBytesAmount], nil
}

func (c *Concater) getCurrentSource() *os.File {
	if c.nextSourceIndex >= len(c.openSourceActions) {
		return nil
	}

	if c.currentSource == nil {
		openSourceAction := c.openSourceActions[c.nextSourceIndex]
		c.currentSource = openSourceAction()
		c.nextSourceIndex++
	}

	return c.currentSource
}

func (c *Concater) closeCurrentSource() {
	if c.currentSource == nil {
		return
	}

	c.currentSource.Close()
	c.currentSource = nil
}

func (c *Concater) write(buffer []byte) {
	_, err := c.target.Write(buffer)
	if err != nil {
		panic(err)
	}
}
