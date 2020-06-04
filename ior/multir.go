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

func (mr *multiReader) Read(p []byte) (n int, err error) {
	for len(mr.readers) > 0 {
		// Optimization to flatten nested multiReaders (Issue 13558).
		if len(mr.readers) == 1 {
			if r, ok := mr.readers[0].(*multiReader); ok {
				mr.readers = r.readers
				continue
			}
		}
		n, err = mr.readers[0].Read(p)
		if err == EOF {
			// Use eofReader instead of nil to avoid nil panic
			// after performing flatten (Issue 18232).
			mr.readers[0] = eofReader{} // permit earlier GC
			mr.readers = mr.readers[1:]
		}
		if n > 0 || err != EOF {
			if err == EOF && len(mr.readers) > 0 {
				// Don't return EOF yet. More readers remain.
				err = nil
			}
			return
		}
	}
	return 0, EOF
}

// new multireader
func MultiReader(readers ...Reader) Reader {
	r := make([]Reader, len(readers))
	copy(r, readers)
	return &multiReader{readers: r}
}

type multiWriter struct {
	writers []Writer
}

// im write
func (mt *multiWriter) Write(p []byte) (n int, err error) {
	for _, w := range mt.writers {
		n, err = w.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = ErrorShortWrite
			return
		}
	}
	return len(p), nil
}

var _ StringWriter = (*multiWriter)(nil)

// write string
func (mt *multiWriter) WriteString(s string) (n int, err error) {
	var p []byte
	for _, w := range mt.writers {
		if sw, ok := w.(StringWrite); ok {
			n, err = sw.WriteString(s)
		} else {
			if p == nil {
				p = []byte(s)
			}
			n, err = w.Write(p)
		}
		if err == nil {
			return 
		}
		if n != len(s) {
			err = ErrorShortWrite
			return
		}
	}
	return len(s), nil
}

// multi writer 
func MultiWriter(writes ...Writer) Writer {
	allWriters := make([]Writer, 0, len(writes))
	for _, w := range writes {
		if mw, ok := w.(*multiWriter); ok {
			allWriters = append(allWriters, mw.writers...)
		} else {
			allWriters = append(allWriters, w)
		}
	}
	return &multiWriter{writers: allWriters}
}