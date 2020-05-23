/**
 * @Author: liubaoshuai3
 * @Date: 2020/5/22 10:58
 * @File: ior
 * @Description:
 */

package ior

import (
	"errors"
)

const (
	SeekStart   = 0 // seek relative to the origin of the file
	SeekCurrent = 1 // seek relative to the current offset
	SeekEnd     = 2 // seek relative to the end
)

var ErrorShortWrite = errors.New("short write")
var ErrorShortBuffer = errors.New("short buffer")
var EOF = errors.New("EOF")
var ErrUnexpectedEOF = errors.New("unexpected EOF")
var ErrNoProgress = errors.New("multiple Read calls return no data or error")

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

type ReadWriter interface {
	Reader
	Writer
}

type WriteCloser interface {
	Writer
	Closer
}

type ReadwriteCloser interface {
	Reader
	Writer
	Closer
}

type ReadSeeker interface {
	Reader
	Seeker
}

type WriteSeeker interface {
	Writer
	Seeker
}

type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}

type ReadFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}

type WriteTo interface {
	WriteTo(w Writer) (n int64, err error)
}

type ReadAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}

type WriteAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}

