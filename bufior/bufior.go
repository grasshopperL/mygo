/**
 * @Author: liubaoshuai3
 * @Date: 2020/6/4 21:13
 * @File: bufior
 * @Description:
 */

package bufior

import (
	"bytes"
	"errors"
	"io"
	"math"
	"unicode/utf8"
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

func (b *Reader) Discard(n int) (discard int, err error) {
	if n < 0 {
		return 0, ErrNegativeCount
	}
	if n == 0 {
		return
	}
	remain := n
	for {
		skip := b.Buffered()
		if skip == 0 {
			b.fill()
			skip = b.Buffered()
		}
		if skip > remain {
			skip = remain
		}
		b.r += skip
		remain -= skip
		if remain == 0 {
			return n, nil
		}
		if b.err != nil {
			return n - remain, b.readErr()
		}
	}
}

// hwo much can be read again
func (b *Reader) Buffered() int {
	return b.w - b.r
}

func (b *Reader) Read(p []byte) (n int, err error) {
	n = len(p)
	if n == 0 {
		if b.Buffered() > 0 {
			return 0, nil
		}
		return 0, b.readErr()
	}
	if b.r == b.w {
		if b.err != nil {
			return 0, b.readErr()
		}
		if len(p) >= len(b.buf) {
			n, b.err = b.rd.Read(p) // what's the meaning of like this line of code
			if n < 0 {
				panic(ErrNegativeCount)
			}
			if n > 0 {
				b.lastByte = int(p[n - 1])
				b.lastRuneSize = -1
			}
			return n, b.readErr()
		}
		b.r = 0
		b.w = 0
		n, b.err = b.rd.Read(b.buf)
		if n < 0 {
			panic(errNegativeRead)
		}
		if n == 0 {
			return 0, b.readErr()
		}
		b.w += n
	}
	n = copy(p, b.buf[b.r:b.w])
	b.r += n
	b.lastByte = int(b.buf[b.r-1])
	b.lastRuneSize = -1
	return n, nil
}

// read one byte
func (b *Reader) ReadByte() (byte, error) {
	b.lastRuneSize = -1
	for b.r == b.w {
		if b.readErr() != nil {
			return 0, b.readErr()
		}
		b.fill()
	}
	c := b.buf[b.r]
	b.r++
	b.lastByte = int(c)
	return c, nil
}

// unread one byte
func (b *Reader) UnReadByte() error {
	if b.lastByte < 0 || b.w > 0 && b.r == 0 {
		return ErrInvalidUnreadByte
	}
	if b.r > 0 {
		b.r--
	} else {
		b.w = 1
	}
	b.buf[b.r] = byte(b.lastByte)
	b.lastByte = -1
	b.lastRuneSize = -1
	return nil
}

// read rune and it's too difficult to understand
func (b *Reader) ReadRune() (r rune, size int, err error) {
	for b.r + utf8.UTFMax > b.w && !utf8.FullRune(b.buf[b.r:b.w]) && b.err == nil && b.w-b.r < len(b.buf) {
		b.fill()
	}
	b.lastRuneSize = -1
	if b.r == b.w {
		return 0, 0, b.readErr()
	}
	r, size = rune(b.buf[b.r]), 1
	if r >= utf8.RuneSelf {
		r, size = utf8.DecodeRune(b.buf[b.r:b.w])
	}
	b.r += size
	b.lastByte = int(b.buf[b.r-1])
	b.lastRuneSize = size
	return r, size, nil
}

// unread a rune
func (b *Reader) UnreadRune() error {
	if b.lastRuneSize < 0 || b.r < b.lastRuneSize {
		return ErrInvalidUnreadRune
	}
	b.r -= b.lastRuneSize
	b.lastByte = -1
	b.lastRuneSize = -1
	return nil
}

// ?
func (b *Reader) Buffered() int {
	return b.w - b.r
}

// too difficult to understand
func (b *Reader) ReadSlice(delim byte) (line []byte, err error) {
	s := 0
	for {
		if i := bytes.IndexByte(b.buf[b.r+s:b.w], delim); i >= 0 {
			i += s
			line = b.buf[b.r : b.r+i+1]
			b.r += i + 1
			break
		}
		if b.err != nil {
			line = b.buf[b.r:b.w]
			b.r = b.w
			err = b.readErr()
			break
		}
		if b.Buffered() >= len(b.buf) {
			b.r = b.w
			line = b.buf
			err = ErrBufferFull
			break
		}
		s = b.w - b.r
		b.fill()
	}
	if i := len(line) - 1; i >= 0 {
		b.lastByte = int(line[i])
		b.lastRuneSize = -1
	}

	return
}

// difficult
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error) {
	line, err = b.ReadSlice('\n')
	if err == ErrBufferFull {
		if len(line) > 0 && line[len(line)-1] == '\r' {
			if b.r == 0 {
				panic("bufio: tried to rewind past start of buffer")
			}
			b.r--
			line = line[:len(line)-1]
		}
		return line, true, nil
	}
	if len(line) == 0 {
		if err != nil {
			line = nil
		}
		return
	}
	err = nil

	if line[len(line)-1] == '\n' {
		drop := 1
		if len(line) > 1 && line[len(line)-2] == '\r' {
			drop = 2
		}
		line = line[:len(line)-drop]
	}
	return
}

// difficult
func (b *Reader) collectFragments(delim byte) (fullBuffers [][]byte, finalFragment []byte, totalLen int, err error) {
	var frag []byte
	for {
		var e error
		frag, e = b.ReadSlice(delim)
		if e == nil {
			break
		}
		if e != ErrBufferFull {
			err = e
			break
		}
		buf := make([]byte, len(frag))
		copy(buf, frag)
		fullBuffers = append(fullBuffers, buf)
		totalLen += len(buf)
	}
	totalLen += len(frag)
	return
}

func (b *Reader) ReadBytes(delim byte) ([]byte, error) {
	full, frag, n, err := b.collectFragments(delim)
	buf := make([]byte, n)
	n = 0
	for i := range full {
		n += copy(buf[n:], full[i])
	}
	copy(buf[n:], frag)
	return buf, err
}