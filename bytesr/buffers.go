/**
 * @Author: liubaoshuai3
 * @Date: 2020/8/24 18:42
 * @File: buffers
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

var ErrTooLarge = errors.New("bytesr.buffers: too large")
var errNegativeRead = errors.New("bytesr.buffers: reader returned negative count from read")
const maxInt = int(^uint(0) >> 1)

func (b *Buffer) Bytes() []byte {
	return b.buf[b.off:]
}

func (b *Buffer) String() string {
	if b.buf == nil {
		return "<nil>"
	}
	return string(b.buf[b.off:])
}

func (b *Buffer) empty() bool {
	return len(b.buf) <= b.off
}

func (b *Buffer) Len() int {
	return len(b.buf) - b.off
}

func (b *Buffer) Cap() int {
	return cap(b.buf)
}

func (b *Buffer) Truncate(n int) {
	if n == 0 {
		b.Reset()
		return
	}
	b.lastRead = opInvalid
	if n < 0 || n > b.Len() {
		panic("bytesr.buffer: truncation out of range")
	}
	b.buf = b.buf[:b.off+n]
}

