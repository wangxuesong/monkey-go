package lexer

import (
	"github.com/stretchr/testify/assert"
	"github.com/wangxuesong/monkey-go/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectType    token.TokenType
		expectLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
	}

	l := New(input)
	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tt.expectType, tok.Type)
		assert.Equal(t, tt.expectLiteral, tok.Literal)
	}
}
