/**
 * @Author: liubaoshuai3
 * @Date: 2020/6/4 21:13
 * @File: bufior
 * @Description:
 */

package bufior

import (
	"errors"
	"io"
)

const defaultBufSize = 4096

var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)

type Reader struct {
	buf []byte
	rd io.Reader
	r, w int
	err error
	lastByte int
	lastRuneSize int
}

const minReadBufferSize = 16
const maxConsecutiveEmptyReads = 100

// how to ignore warn
func NewReaderSize(rd io.Reader, size int) *Reader {
	if b, ok := rd.(*Reader); ok && len(b.buf) > size {
		return b
	}
	if size < minReadBufferSize {
		size = minReadBufferSize
	}
	r := new(Reader)
	r.reset(make([]byte, size), rd)
	return r
}

// new buf reader by default size
func NewBufReader(rd io.Reader) *Reader {
	return NewReaderSize(rd, defaultBufSize)
}

// return the underlying buf size
func (b *Reader) Size() int {
	return len(b.buf)
}

// reset reader
func (b *Reader) Reset(rd io.Reader) {
	b.reset(b.buf, rd)
}

// reset Reader and change reader with r
func (b *Reader) reset(buf []byte, r io.Reader) *Reader {
	return &Reader{
		buf:          buf,
		rd:           r,
		r:            0,
		w:            0,
		err:          nil,
		lastByte:     -1,
		lastRuneSize: -1,
	}
}
