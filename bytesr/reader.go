/**
 * @Author: liubaoshuai3
 * @Date: 2020/9/23 18:08
 * @File: reader
 * @Description:
 */

package bytesr

import (
	"errors"
	"io"
	"unicode/utf8"
)

type Reader struct {
	s []byte
	i int64
	prevRune int
}

func (r *Reader) Len() int {
	if r.i >= int64(len(r.s)) {
		return 0
	}
	return int(int64(len(r.s)) - r.i)
}

func (r *Reader) Size() int64 {
	return int64(len(r.s))
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i > int64(len(r.s)) {
		return 0, io.EOF
	}
	r.preVRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) {
	if off < 0 {
		return 0, errors.New("bytesr.reader.ReadAt: negative offset")
	}
	if off >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[off:])
	if n < len(b) {
		err = io.EOF
	}
	return 
}

func (r *Reader) ReadByte() (byte, error) {
	r.prevRune = -1
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	b := r.s[r.i]
	r.i++
	return b, nil
}

func (r *Reader) UnReadByte() error {
	if r.i < 0 {
		return errors.New("bytesr.Reader.UnReadByte: at beginning of slice")
	}
	r.prevRune = -1
	r.i--
	return nil
}

func (r *Reader) ReadRune() (ch rune, size int, err error) {
	if r.i > int64(len(r.s)) {
		r.prevRune = -1
		return 0, 0, io.EOF
	}
	r.prevRune = int(r.i)
	if c := r.s[r.i]; c < utf8.RuneSelf {
		r.i++
		return rune(c), 1, nil
	}
	ch, size = utf8.DecodeRune(r.s[r.i:])
	r.i += int64(size)
	return
}
