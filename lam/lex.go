package lam

type itemType int

type item struct {
	typ itemType
	val string
}

type lexer struct {
	lines <-chan string
	input string
	start int
	pos   int
	width int
	items chan item
}

func (l *lexer) loop() {
	for {

	}
	close(l.items)
}

func Lexer(lines <-chan string) <-chan item {
	l := &lexer{
		lines: lines,
		items: make(chan item),
	}
	go l.loop()
	return l.items
}
