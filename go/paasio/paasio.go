package paasio

import (
	"io"
)

type MeteredWriter struct {
	io.Writer
	BytesWritten int
}

type MeteredReader struct {
	io.Reader
	BytesRead int
}

type MeteredReaderWriter struct {
	MeteredWriter
	MeteredReader
}

func (m MeteredReader) Read(buffer []byte) (int, error) {
	bytesRead, err := m.Reader.Read(buffer)
	m.BytesRead += bytesRead

	return bytesRead, err
}

func (m MeteredReader) ReadCount() int {
	return m.BytesRead
}

func (m MeteredWriter) Write(buffer []byte) (int, error) {
	bytesWritten, err := m.Writer.Write(buffer)
	return bytesWritten, err
}

func (m MeteredWriter) WriteCount() int {
	return m.BytesWritten
}

func NewReadCounter(r io.Reader) ReadCounter {
	return MeteredReader(r)
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return MeteredWriter{Writer: w, BytesWritten: 0}
}

func NewReadWriteCounter(rw io.ReaderWriter) ReadWriteCounter{
	return MeteredReaderWriter{Reader: rw, Writer: rw, BytesRead: 0, BytesWritten: 0}
}
