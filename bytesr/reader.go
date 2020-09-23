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
)

type Reader struct {
	s []byte
	i int64
	preVRune int
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
