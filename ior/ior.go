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

type ByteReader interface {
	ReadByte() (byte, error)
}

type ReadScanner interface {
	ByteReader
	UnreadByte() error
}

type ByteWriter interface {
	WriteByte(c byte) error
}

type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}

type RuneScanner interface {
	RuneReader
	UnreadRune() error
}

type StringWrite interface {
	WriteString(s string) (n int, err error)
}

// write string to writer
func WriteString(w Writer, s string) (n int, err error) {
	if sw, ok := w.(StringWrite); ok {
		return sw.WriteString(s)
	}
	return w.Write([]byte(s))
}

// read at least min bytes
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error) {
	if len(buf) < min {
		return 0, ErrorShortBuffer
	}
	for n < min && err == nil {
		var nn int
		nn, err = r.Read(buf[n:])
		n += nn
	}
	if n >= min {
		err = nil
	} else if n > 0 && err == EOF {
		err = ErrUnexpectedEOF
	}
	return 
}