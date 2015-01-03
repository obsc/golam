package lam

import (
	"bufio"
	"io"
)

type Scanner struct {
	buf *bufio.Reader
}

func (s *Scanner) run() {
	for {

	}
}

func ScanReader(rd io.Reader) *Scanner {
	s := &Scanner{
		buf: bufio.NewReader(rd),
	}
	go s.run()
	return s
}
