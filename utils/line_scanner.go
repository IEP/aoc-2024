package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	ReadBufferSize = 100
)

type LineReader struct {
	f *os.File
	s *bufio.Scanner

	closed bool
}

func NewLineReader(filename string) (*LineReader, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	s := bufio.NewScanner(f)

	return &LineReader{
		f: f,
		s: s,
	}, nil
}

func (r *LineReader) Line() (string, error) {
	if !r.s.Scan() {
		r.f.Close()
		r.closed = true
	}

	if r.closed {
		return "", io.EOF
	}

	return r.s.Text(), nil
}
