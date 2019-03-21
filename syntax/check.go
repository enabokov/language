package syntax

import (
	"fmt"

	"github.com/enabokov/language/lexis"
)

var Precedence = map[string]int{
	"=":  1,
	"||": 2,
	"&&": 3,
	"<":  7, ">": 7, "<=": 7, ">=": 7, "==": 7, "!=": 7,
	"+": 10, "-": 10,
	"*": 20, "/": 20, "%": 20,
}

func IsPunctuation(input lexis.TokenStream, lexeme string) *lexis.Token {
	token := input.Peek()

	if token.Class != lexis.ClassPunctuation {
		return nil
	}

	if lexeme == "" || token.Value != lexeme {
		return nil
	}

	return token
}

func IsKeyword(input lexis.TokenStream, lexeme string) *lexis.Token {
	token := input.Peek()

	if token.Class != lexis.ClassKeyword {
		return nil
	}

	if lexeme == "" || token.Value != lexeme {
		return nil
	}

	return token
}

func IsOperator(input lexis.TokenStream, lexeme string) *lexis.Token {
	token := input.Peek()

	if token.Class != lexis.ClassOperator {
		return nil
	}

	if lexeme == "" || token.Value != lexeme {
		return nil
	}

	return token
}

func SkipPunctuation(input lexis.TokenStream, lexeme string) {
	if IsPunctuation(input, lexeme) != nil {
		input.Next()
	} else {
		input.Croak(fmt.Sprintf("Expecting punctuation: %s", lexeme))
	}
}

func SkipKeyword(input lexis.TokenStream, lexeme string) {
	if IsKeyword(input, lexeme) != nil {
		input.Next()
	} else {
		input.Croak(fmt.Sprintf("Expecting keyword: %s", lexeme))
	}
}

func SkipOperator(input lexis.TokenStream, lexeme string) {
	if IsOperator(input, lexeme) != nil {
		input.Next()
	} else {
		input.Croak(fmt.Sprintf("Expecting operator: %s", lexeme))
	}
}

func Unexpected(input lexis.TokenStream) {
	input.Croak(fmt.Sprintf("Unexpected token: %s", input.Peek()))
}
