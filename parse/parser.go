package parse

type Token struct {
	Line, Kind int
	Value      string
}
type Parser struct {
	lexer  *Lexer
	tokens []*Token
	curIdx int
	eof    bool
}

func NewToken(line, kind int, value string) *Token {
	return &Token{
		Line: line, Kind: kind, Value: value,
	}
}
func NewParser(l *Lexer) *Parser {
	return &Parser{
		lexer:  l,
		tokens: make([]*Token, 0),
		curIdx: -1,
		eof:    false,
	}
}
