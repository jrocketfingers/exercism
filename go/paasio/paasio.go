package paasio

import (
	"io"
)

type MeteredWriter struct {
	io.Writer
	n    int64
	nops int
}

type MeteredReader struct {
	io.Reader
	n    int64
	nops int
}

type MeteredReaderWriter struct {
	MeteredWriter
	MeteredReader
}

func (m MeteredReader) Read(buffer []byte) (bytesRead int, err error) {
	bytesRead, err = m.Reader.Read(buffer)
	m.n += int64(bytesRead)
	m.nops += 1
	return
}

func (m MeteredReader) ReadCount() (n int64, nops int) {
	n = m.n
	nops = m.nops
	return
}

func (m MeteredWriter) Write(buffer []byte) (bytesWritten int, err error) {
	bytesWritten, err = m.Writer.Write(buffer)
	m.n += int64(bytesWritten)
	m.nops += 1
	return
}

func (m MeteredWriter) WriteCount() (n int64, nops int) {
	n = m.n
	nops = m.nops
	return
}

func NewReadCounter(r io.Reader) ReadCounter {
	return MeteredReader{Reader: r}
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return MeteredWriter{Writer: w}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return MeteredReaderWriter{MeteredReader: MeteredReader{Reader: rw}, MeteredWriter: MeteredWriter{Writer: rw}}
}
