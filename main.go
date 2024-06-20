package main

import (
	"fmt"
	"os"

	"chenqijia1234.github.io/kpp/parse"
)

func main() {
	file_name := os.Args[1]
	t, err := os.ReadFile(file_name)
	if err != nil {
		panic("OS Error")
	}
	lexer := parse.NewLexer(string(t), file_name)
	line, kind, tok := lexer.NextToken()
	for kind != parse.TOKEN_EOF {
		fmt.Println(line, parse.KindToCategory(kind), tok)
		line, kind, tok = lexer.NextToken()
	}
}
