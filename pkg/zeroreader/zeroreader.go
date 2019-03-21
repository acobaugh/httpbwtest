package zeroreader

import "io"

type ZeroReader struct {
	size int
}

func NewZeroReader(size int) *ZeroReader {
	return &ZeroReader{size: size}
}

func (a *ZeroReader) Read(p []byte) (int, error) {
	return 0, io.EOF
}
