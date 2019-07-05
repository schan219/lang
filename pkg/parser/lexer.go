package lexer

type TokenType int
type StateType int

const (
	TokenEmpty TokenType = iota

	TokenLParen
	TokenRParen

	TokenDot
	TokenString
	TokenSymbol
)

struct Token {
	// The type of token we've parsed
	Type TokenType
	// The string value of the token
	Value string

	// The line and column for number
	Line int
	Col int
}


const (
	StateType = iota
)


func CommentFSM() {
	switch currentChar {
	case "/":	
	}
}