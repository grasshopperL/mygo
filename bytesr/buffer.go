/**
 * @Author: liubaoshuai3
 * @Date: 2020/5/14 10:18
 * @File: buffer
 * @Description:
 */

package bytesr

import "errors"

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