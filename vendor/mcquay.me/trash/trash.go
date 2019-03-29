// Package trash provides simple readers with meaningless output.
package trash

import (
	"io"
	"math/rand"
	"time"
)

func init() {
	Reader = &reader{0xca}
	Zeros = &reader{0x00}
	Fs = &reader{0xff}
	HiLo = &reader{0xaa}
	LoHi = &reader{0x55}
	Random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Reader provides a steady stream of trash (non-random bytes) when read from
var Reader io.Reader

// Zeros provides a steady stream of 0x00
var Zeros io.Reader

// Fs provides a steady stream of 0xff
var Fs io.Reader

// HiLo provides a steady stream of 0xaa
var HiLo io.Reader

// LoHi provides a steady stream of 0x55
var LoHi io.Reader

var Random io.Reader

// TimeoutReader returns a reader that returns io.EOF after dur.
func TimeoutReader(dur time.Duration) io.Reader {
	return &timeoutReader{timeout: time.Now().Add(dur)}
}

type reader struct {
	pattern byte
}

func (r *reader) Read(p []byte) (int, error) {
	c := 0
	var err error
	for i := 0; i < len(p); i++ {
		c++
		p[i] = r.pattern
	}
	return c, err
}

type timeoutReader struct {
	timeout time.Time
}

func (tor *timeoutReader) Read(p []byte) (int, error) {
	c := 0
	var err error
	for i := 0; i < len(p); i++ {
		c++
		p[i] = 0xca
	}
	if time.Now().After(tor.timeout) {
		err = io.EOF
	}
	return c, err
}
