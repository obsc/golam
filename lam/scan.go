package lam

import (
	"bufio"
	"io"
)

type scanner struct {
	buf   *bufio.Reader
	lines chan string
}

func (s *scanner) loop() {
	for {
		line, err := s.buf.ReadString('\n')
		s.lines <- line
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	close(s.lines)
}

func Scanner(rd io.Reader) <-chan string {
	s := &scanner{
		buf:   bufio.NewReader(rd),
		lines: make(chan string),
	}
	go s.loop()
	return s.lines
}
