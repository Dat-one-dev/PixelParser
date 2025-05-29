package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	ILLEGAL

	// Add these new token types
	NUMBER       // 123, 10, etc.
	STRING       // "Hello, Atomicity!"

	// Comments
	COMMENT_LINE  // // ...
	COMMENT_BLOCK // /* ... */

	// Identifiers & Literals
	IDENTIFIER   // variable names (e.g., x, y, z)
	PRIMITIVE    // number, bool, or your single-value type
	RGBA_LITERAL // [255,0,0,255]
	HEX_LITERAL  // #FF00FF

	// Keywords / Built-ins
	PRINT     // print
	DISPLAY   // display
	PIXEL     // pixel
	PUSH      // push
	RGBA_FUNC // rgba()
	HEX_FUNC  // hex()

	// Operators / Symbols
	ASSIGN   // =
	COMMA    // ,
	COLON    // :
	DOT      // .
	LPAREN   // (
	RPAREN   // )
	LBRACKET // [
	RBRACKET // ]
)

type Token struct {
	Kind  TokenKind
	Value string
}

func (token Token) isOneOfMany(expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
		if expected == token.Kind {
			return true
		}
	}
	return false
}

func (token Token) Debug() {
	if token.isOneOfMany(IDENTIFIER, PRIMITIVE, RGBA_LITERAL, HEX_LITERAL, NUMBER, STRING) {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(token.Kind))
	}
}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case ILLEGAL:
		return "illegal"
	case COMMENT_LINE:
		return "comment_line"
	case COMMENT_BLOCK:
		return "comment_block"
	case IDENTIFIER:
		return "identifier"
	case PRIMITIVE:
		return "primitive"
	case RGBA_LITERAL:
		return "rgba_literal"
	case HEX_LITERAL:
		return "hex_literal"
	case PRINT:
		return "print"
	case DISPLAY:
		return "display"
	case PIXEL:
		return "pixel"
	case PUSH:
		return "push"
	case RGBA_FUNC:
		return "rgba_func"
	case HEX_FUNC:
		return "hex_func"
	case ASSIGN:
		return "assign"
	case COMMA:
		return "comma"
	case COLON:
		return "colon"
	case DOT:
		return "dot"
	case LPAREN:
		return "lparen"
	case RPAREN:
		return "rparen"
	case LBRACKET:
		return "lbracket"
	case RBRACKET:
		return "rbracket"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}
