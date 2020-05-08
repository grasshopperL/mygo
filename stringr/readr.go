/**
 * @Author: liubaoshuai3
 * @Date: 2020/5/7 10:49
 * @File: readr
 * @Description:
 */

package stringr

import (
	"errors"
	"io"
)

type Readr struct {
	s string
	i int64
	prevRune int
}

// the length of unread, why is int
func (r *Readr) Len() int {
	if r.i >= int64(len(r.s)) {
		return 0
	}
	return int(int64(len(r.s)) - r.i)
}

// why is int64
func (r *Readr) Size() int64 {
	return int64(len(r.s))
}

// read the remain
func (r *Readr) Read(b []byte) (n int, e error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

// some ???
func (r *Readr) ReadAt(b []byte, off int) (n int, err error) {
	if off < 0 {
		return 0, errors.New("stringr.Readr.ReadAt: negative offset")
	}
	if off > len(r.s) {
		return 0, io.EOF
	}
	n = copy(b, r.s[off:])
	if n < len(b) {
		err = io.EOF
	}
	return
}