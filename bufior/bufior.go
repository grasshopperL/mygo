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

// read a chunk to buffer
// why tried limit times as for
func (b *Reader) fill() {
	if b.r > 0 {
		copy(b.buf, b.buf[b.r:b.w])
		b.w -= b.r
		b.r = 0
	}
	if b.w >= len(b.buf) {
		panic("bufior.bufior.fill: tried to write full buffer")
	}
	for i := maxConsecutiveEmptyReads; i > 0; i-- {
		n, err := b.rd.Read(b.buf[b.w:])
		if n < 0 {
			panic(ErrNegativeCount)
		}
		b.w += n
		if err != nil {
			b.err = err
			return
		}
		if n > 0 {
			return
		}
	}
	b.err = io.ErrNoProgress
}

// get the read err and why it's will reset the err
func (b *Reader) readErr() error {
	err := b.err
	b.err = nil
	return err
}

// peek where can use it
func (b *Reader) Peek(n int) ([]byte, error) {
	if n < 0 {
		return nil, ErrNegativeCount
	}
	b.lastRuneSize = -1
	b.lastByte = -1
	for b.w - b.r < n && b.w - b.r < len(b.buf) && b.err == nil {
		b.fill()
	}
	if n > len(b.buf) {
		return b.buf[b.r:b.w], ErrBufferFull
	}
	var err error
	if available := b.w - b.r; available < n {
		n = available
		err = b.readErr()
		if err == nil {
			err = ErrBufferFull
		}
	}
	return b.buf[b.r:b.r + n], err
}