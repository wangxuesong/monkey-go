package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

// 标识符
const (
	IDENT = "IDENT"
	INT   = "INT"
)

// 操作符
const (
	ASSIGN = "="
	PLUS   = "+"
)

// 分隔符
const (
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
)

// 关键字
const (
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
