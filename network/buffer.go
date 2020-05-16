package network

import (
	"errors"
	"io"
)

type Buffer struct {
	reader io.Reader
	buf    []byte
	start  int
	end    int
}

func newBuffer(reader io.Reader, len int) Buffer {
	return Buffer{
		reader: reader,
		buf:    make([]byte, len),
		start:  0,
		end:    0,
	}
}

func (b *Buffer) len() int {
	return b.end - b.start
}

func (b *Buffer) readFromReader() error {
	b.grow()
	n, err := b.reader.Read(b.buf[b.end:])
	if err != nil {
		return err
	}
	b.end += n
	return nil
}

func (b *Buffer) grow() {
	if b.start == 0 {
		return
	}
	copy(b.buf[0:], b.buf[b.start:b.end])
	b.end = b.end - b.start
	b.start = 0
}

func (b *Buffer) seek(offset, limit int) ([]byte, error) {
	if b.len() < offset+limit {
		return nil, errors.New("not enough")
	}
	return b.buf[b.start+offset : b.start+offset+limit], nil
}

func (b *Buffer) read(offset, limit int) ([]byte, error) {
	if b.len() < limit-offset {
		return nil, errors.New("not enough")
	}
	b.start += offset
	buf := b.buf[b.start : b.start+limit]
	b.start += limit
	return buf, nil
}
