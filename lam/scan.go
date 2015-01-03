package lam

import (
	"bufio"
	"io"
	"strings"
)

type Scanner struct {
	buf   *bufio.Reader
	Lines chan string
}

func (s *Scanner) run() {
	for {
		line, err := s.buf.ReadString('\n')
		if err == io.EOF {
			s.Lines <- strings.TrimRight(line, "\n")
			break
		} else if err == nil {
			s.Lines <- line
		} else {
			panic(err)
		}
	}
	close(s.Lines)
}

func ScanReader(rd io.Reader) *Scanner {
	s := &Scanner{
		buf:   bufio.NewReader(rd),
		Lines: make(chan string),
	}
	go s.run()
	return s
}
