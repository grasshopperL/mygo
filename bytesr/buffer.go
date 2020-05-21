/**
 * @Author: liubaoshuai3
 * @Date: 2020/5/14 10:18
 * @File: buffer
 * @Description:
 */

package bytesr

import (
	"errors"
	"io"
	"unicode/utf8"
)

const smallBufferSize = 64

type Buffer struct {
	buf []byte
	off int
	lastRead readOp
}

type readOp int8

const (
	opRead      readOp = -1 // Any other read operation.
	opInvalid   readOp = 0  // Non-read operation.
	opReadRune1 readOp = 1  // Read rune of size 1.
	opReadRune2 readOp = 2  // Read rune of size 2.
	opReadRune3 readOp = 3  // Read rune of size 3.
	opReadRune4 readOp = 4  // Read rune of size 4.
)

var ErrorTooLarge = errors.New("bytesr:buffer: too large")
var errNegativeRead = errors.New("bytesr.buffer: reader returned negative count from Read")

const maxInt = int(^uint(0) >> 1)

// return un read
func (b *Buffer) Bytes() []byte {
	return b.buf[b.off:]
}

// return unread
func (b *Buffer) String() string {
	if b.buf == nil {
		return "<nil>"
	}
	return string(b.buf[b.off:])
}

// buffer is empty
func (b *Buffer) empty() bool {
	return len(b.buf) < b.off
}

// length of unread b.Len() == len(b.Bytes())
func (b *Buffer) Len() int {
	return len(b.buf) - b.off
}

// the cap of b.buf
func (b *Buffer) Cap() int {
	return cap(b.buf)
}

// keep first n
func (b *Buffer) Truncate(n int) {
	if n == 0 {
		b.Reset()
		return
	}
	b.lastRead = opInvalid
	if n < 0 || n > len(b.buf) {
		panic("bytes.Buffer:truncate out off range")
	}
	b.buf = b.buf[:b.off + n]
}

// reset the buff
func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
	b.lastRead = opInvalid
	b.off = 0
}

// seems like grow
func (b *Buffer) tyrGrowByReslice(n int) (int, bool) {
	if l := len(b.buf); n <= cap(b.buf) - l {
		b.buf = b.buf[:b.off + l]
		return l, true
	}
	return 0, false
}

// real grow, return index where bytes should be written
func (b *Buffer) grow(n int) int {
	m := b.Len()
	if m == 0 && b.off != 0 {
		b.Reset()
	}
	if i, ok := b.tyrGrowByReslice(n); ok {
		return i
	}
	if b.buf == nil && n <= smallBufferSize {
		b.buf = make([]byte, n, smallBufferSize)
		return 0
	}
	c := cap(b.buf)
	if n <= c/2 -m {
		copy(b.buf, b.buf[:])
	} else if c > maxInt -c -n {
		panic(ErrorTooLarge)
	} else {
		buf := makeSlice(2 * c + n)
		copy(buf, b.buf[b.off:])
		b.buf = buf
	}
	b.off = 0
	b.buf = b.buf[:m + n]
	return m
}

// grow
func (b *Buffer) Grow(n int) {
	if n < 0 {
		panic("bytesr.buffer.Grow: negative count")
	}
	m := b.grow(n)
	b.buf = b.buf[:m]
}

// write appends the content of p to buffer
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.lastRead = opInvalid
	m, ok := b.tyrGrowByReslice(len(p))
	if !ok {
		m = b.grow(len(p))
	}
	return copy(b.buf[m:], p), nil
}

//
func (b *Buffer) WriteString(s string) (n int, err error) {
	b.lastRead = opInvalid
	m, ok := b.tyrGrowByReslice(len(s))
	if !ok {
		m = b.grow(len(s))
	}
	return copy(b.buf[m:], s), nil
}

const MinRead = 512

func (b *Buffer) ReadForm(r io.Reader) (n int64, err error) {
	b.lastRead = opInvalid
	for {
		i := b.grow(MinRead)
		b.buf = b.buf[:i]
		m, e := r.Read(b.buf[i:cap(b.buf)])
		if m < 0 {
			panic(errNegativeRead)
		}
		b.buf = b.buf[:i + m]
		n += int64(m)
		if e == io.EOF {
			return n, nil
		}
		if e != nil {
			return n, e
		}
	}
}

// makeSlice allocates a slice of size n.
func makeSlice(n int) []byte {
	defer func() {
		if recover() != nil {
			panic(ErrorTooLarge)
		}
	}()
	return make([]byte, n)
}

// write one byte to buffer
func (b *Buffer) WriteByte(c byte) error {
	b.lastRead = opInvalid
	m, ok := b.tyrGrowByReslice(1)
	if !ok {
		m = b.grow(1)
	}
	b.buf[m] = c
	return nil
}

func (b *Buffer) WriteRune(r rune) (n int, err error) {
	if r < utf8.RuneSelf {
		b.WriteByte(byte(r))
		return 1, nil
	}
	b.lastRead = opInvalid
	m, ok := b.tyrGrowByReslice(utf8.UTFMax)
	if !ok {
		m = b.grow(utf8.UTFMax)
	}
	n = utf8.EncodeRune(b.buf[m:m+utf8.UTFMax], r)
	b.buf = b.buf[:m+n]
	return n, nil
}

// read len(p) or all b.buf
func (b *Buffer) Read(p []byte) (n int, err error) {
	b.lastRead = opInvalid
	if b.empty() {
		b.Reset()
		if len(p) == 0 {
			return 0, nil
		}
		return 0, io.EOF
	}
	n = copy(p, b.buf[b.off:])
	b.off += n
	if n > 0 {
		b.lastRead = opRead
	}
	return n, nil
}

// the next
func (b *Buffer) Next(n int) []byte {
	b.lastRead = opRead
	m := b.Len()
	if n > m {
		n = m
	}
	data := b.buf[b.off:b.off + n]
	b.off += n
	if n > 0 {
		b.lastRead = opRead
	}
	return data
}

// read one byte from buffer
func (b *Buffer) ReadByte() (byte, error) {
	if b.empty() {
		b.Reset()
		return 0, io.EOF
	}
	c := b.buf[b.off]
	b.off++
	b.lastRead = opRead
	return c, nil
}

// read rune from buffer
func (b *Buffer) ReadRune() (r rune, size int, err error) {
	if b.empty() {
		b.Reset()
		return 0, 0, io.EOF
	}
	c := b.buf[b.off]
	if c < utf8.RuneSelf {
		b.off++
		b.lastRead = opReadRune1
		return rune(c), 1, nil
	}
	r, n := utf8.DecodeRune(b.buf[b.off:])
	b.off += n
	b.lastRead = readOp(n)
	return r, n, nil
}

// unread the last rune return by readrune
func (b *Buffer) UnReadRune() error {
	if b.lastRead <= opRead {
		return errors.New("bytesr.buffer.UnReadRune: previous operation is not a successful read operation");
	}
	if b.off > int(b.lastRead) {
		b.off -= int(b.lastRead)
	}
	b.lastRead = opInvalid
	return nil
}

var errUnreadByte = errors.New("bytes.Buffer: UnreadByte: previous operation was not a successful read")

func (b *Buffer) UnReadByte() error {
	if b.lastRead == opInvalid {
		return errUnreadByte
	}
	b.lastRead = opInvalid
	if b.off > 0 {
		b.off--
	}
	return nil
}

// read line ???
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) {
	slice, err := b.readSlice(delim)
	line = append(line, slice...)
	return line, err
}

// I don't know why 
func (b *Buffer) readSlice(delim byte) (line []byte, err error) {
	i := IndexByte(b.buf[b.off:], delim)
	end := b.off + i + 1
	if i < 0 {
		end = len(b.buf)
		err = io.EOF
	}
	line = b.buf[b.off:end]
	b.off = end
	b.lastRead = opRead
	return line, err
}