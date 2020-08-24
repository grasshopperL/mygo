/**
 * @Author: liubaoshuai3
 * @Date: 2020/8/24 18:42
 * @File: buffers
 * @Description:
 */

package bytesr

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

