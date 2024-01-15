package lexer

type Lexer struct {
	input   string
	pos     int  // cur pos in input points to cur char
	readPos int  // cur reading pos in input after cur char
	ch      byte // cur char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

// get the next char and advance the postion
func (l *Lexer) readChar() {
	if l.readPos > len(l.input) { // check if the input is less then the pos
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos += 1 // go up one pos
}
