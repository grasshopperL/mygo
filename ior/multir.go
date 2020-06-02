/**
 * @Author: liubaoshuai3
 * @Date: 2020/6/2 9:52
 * @File: multir
 * @Description:
 */

package ior

type eofReader struct {}

func (e *eofReader) Read([]byte) (int, error) {
	return 0, EOF
}

type multiReader struct {
	readers []Reader
}

func (m *multiReader) Read(p []byte) (n int, err error) {
	
}