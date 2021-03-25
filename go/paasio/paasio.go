package paasio

import (
	"io"
	"sync"
)

type MeteredWriter struct {
	io.Writer
	n          int64
	nops       int
	writeMutex sync.RWMutex
}

type MeteredReader struct {
	io.Reader
	n         int64
	nops      int
	readMutex sync.RWMutex
}

type MeteredReaderWriter struct {
	MeteredWriter
	MeteredReader
}

func (m *MeteredReader) Read(buffer []byte) (bytesRead int, err error) {
	bytesRead, err = m.Reader.Read(buffer)
	m.readMutex.Lock()
	m.n += int64(bytesRead)
	m.nops += 1
	m.readMutex.Unlock()
	return
}

func (m *MeteredReader) ReadCount() (n int64, nops int) {
	m.readMutex.RLock()
	n = m.n
	nops = m.nops
	m.readMutex.RUnlock()
	return
}

func (m *MeteredWriter) Write(buffer []byte) (bytesWritten int, err error) {
	bytesWritten, err = m.Writer.Write(buffer)
	m.writeMutex.Lock()
	m.n += int64(bytesWritten)
	m.nops += 1
	m.writeMutex.Unlock()
	return
}

func (m *MeteredWriter) WriteCount() (n int64, nops int) {
	m.writeMutex.RLock()
	n = m.n
	nops = m.nops
	m.writeMutex.RUnlock()
	return
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &MeteredReader{Reader: r}
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &MeteredWriter{Writer: w}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &MeteredReaderWriter{MeteredReader: MeteredReader{Reader: rw}, MeteredWriter: MeteredWriter{Writer: rw}}
}
