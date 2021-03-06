/**
 * @Author: liubaoshuai3
 * @Date: 2020/5/6 18:24
 * @File: buildr
 * @Description:
 */

package stringr

import "unsafe"

type Buildr struct {
	addr *Buildr
	buf []byte
}

func (b *Buildr) Grow(n int) {
	b.copyCheck()
	if n < 0 {
		panic("stringr.Buildr.Grow: negative param")
	}
	if cap(b.buf) - len(b.buf) < n {
		b.grow(n)
	}
}

func (b *Buildr) grow(n int) {
	buf := make([]byte, len(b.buf), 2 * cap(b.buf) + n)
	copy(buf, b.buf)
	b.buf = buf
}

func (b *Buildr) copyCheck() {
	if b.addr == nil {
		b.addr = (*Buildr)(noescape(unsafe.Pointer(b)))
	}
}

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func (b *Buildr) Reset() {
	b.addr = nil
	b.buf = nil
}

// write some byte
func (b *Buildr) Write(p []byte) (int, error) {
	b.copyCheck()
	b.buf = append(b.buf, p...)
	return len(b.buf), nil
}

// write byte
func (b *Buildr) WriteByte(c byte) error {
	b.copyCheck()
	b.buf = append(b.buf, c)
	return nil
}

// write string
func (b *Buildr) WriteString(s string) (int, error) {
	b.copyCheck()
	b.buf = append(b.buf, s...)
	return len(s), nil
}

