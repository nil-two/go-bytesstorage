// Package bs provides a storage
// to hold the bytes sent over the write of io.writer.
package bs

// BytesStorage is a storage to hold the bytes sent over the write of io.writer.
type BytesStorage struct {
	buf []byte
}

// Write add the bytes in storage.
func (w *BytesStorage) Write(p []byte) (int, error) {
	w.buf = append(w.buf, p...)
	return len(p), nil
}

// Load returns contents of the storage.
func (w *BytesStorage) Load() []byte {
	return w.buf
}

// NewBytesStorage returns a new BytesStorage
func NewBytesStorage() *BytesStorage {
	return &BytesStorage{
		buf: make([]byte, 0),
	}
}
