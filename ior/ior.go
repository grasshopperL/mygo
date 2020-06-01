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

// read len(buf) from reader to buf
func ReadFull(r Reader, buf []byte) (n int, err error) {
	return ReadAtLeast(r, buf, len(buf))
}

// read n byute from src to dst
func CopyN(dst Writer, src Reader, n int64) (written int64, err error) {
	written, err = Copy(dst, LimitReader(src, n))
	if written == n {
		return n, nil
	}
	if written < n && err == nil {
		err = EOF
	}
	return
}

func Copy(dst Writer, src Reader) (written int64,err error) {
	return copyBuffer(dst, src, nil)
}

func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	if len(buf) == 0 && buf != nil {
		panic("ior.CopyBuffer: empty buffer")
	}
	return copyBuffer(dst, src, buf)
}

func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	if wt, ok := dst.(WriteTo); ok {
		return wt.WriteTo(dst)
	}
	if rt, ok := src.(ReadFrom); ok {
		return rt.ReadFrom(src)
	}
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = ErrorShortWrite
				break
			}
		}
		if er != nil {
			if er != EOF {
				err = er
			}
			break
		}
	}
	return 
}

func LimitReader(r Reader, n int64) Reader { return &LimitedReader{r, n} }

type LimitedReader struct {
	R Reader
	N int64 
}

// why call Reader
func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, EOF
	}
	if int64(len(p)) > l.N {
		p = p[:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

type SectionReader struct {
	r ReadAt
	base int64
	off int64
	limit int64
}

// new section reader
func NewSectionReader(r ReadAt, off int64, n int64) *SectionReader {
	return &SectionReader{
		r:     r,
		base:  off,
		off:   off,
		limit: off + n,
	}
}

func (s *SectionReader) Read(p []byte) (n int, err error) {
	if s.limit <= s.off {
		return 0, EOF
	}
	if max := s.limit - s.off; int64(len(p)) > max {
		p = p[:max]
	}
	n, err = s.r.ReadAt(p, s.off)
	s.off += int64(n)
	return
}

var errOffset = errors.New("Seek: invalid set")
var errWhence = errors.New("Seek: invalid whence")

// seek the position
func (s *SectionReader) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	default:
		return 0, errWhence
	case SeekStart:
		offset += s.base
	case SeekCurrent:
		offset += s.off
	case SeekEnd:
		offset += s.limit
	}
	if offset < s.base {
		return 0, errOffset
	}
	s.off = offset
	return offset - s.base, nil
}

// read from one place
func (s *SectionReader) ReadAt(p []byte, offset int64) (n int, err error) {
	if offset < 0 || offset > s.limit - s.off {
		return 0, EOF
	}
	offset += s.base
	if max:= s.limit - offset; max < int64(len(p)) {
		p = p[:max]
		n, err = s.r.ReadAt(p, offset)
		if err == nil {
			err = EOF
		}
		return n, err
	}
	return s.r.ReadAt(p, offset)
}

// the size of section reader
func (s *SectionReader) Size() int64 {
	return s.limit - s.base
}
 // new teeReader
func TeeReader(r Reader, w Writer) Reader {
	return &teeReader{
		r: r,
		w: w,
	}
}

type teeReader struct {
	r Reader
	w Writer
}

func (t *teeReader) Read(p []byte) (n int, err error) {
	n, err = t.r.Read(p)
	if n > 0 {
		if n, err := t.w.Write(p[:n]); err != nil {
			return n, err
		}
	}
	return
}

