package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifier + literal
	IDENT = "IDENT"
	INT   = "INT"

	// operator
	ASSIGN = "="
	PLUS   = "+"

	// delimeter
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keyword
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// func main() {
// 	fmt.Println("vim-go")
// }
