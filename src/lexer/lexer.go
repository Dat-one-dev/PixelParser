package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *lexer, regex *regexp.Regexp)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	pos      int
}

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

func Tokenize(source string) []Token {
	lex := createLexer(source)

	for !lex.at_eof() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())
			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			// Get the first 20 characters of the remaining input for better error context
			preview := lex.remainder()
			if len(preview) > 20 {
				preview = preview[:20] + "..."
			}
			panic(fmt.Sprintf("Lexer::Error -> Unrecognized token near '%s'", preview))
		}
	}

	lex.push(NewToken(EOF, "EOF"))
	return lex.Tokens
}

// === HANDLERS ===

func defaultHandler(kind TokenKind) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		match := regex.FindString(lex.remainder())
		lex.push(NewToken(kind, match))
		lex.advanceN(len(match))
	}
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])
}

func createLexer(source string) *lexer {
	return &lexer{
		source: source,
		Tokens: []Token{},
		patterns: []regexPattern{
			// Skip whitespace and comments
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`//[^\n]*`), skipHandler},
			{regexp.MustCompile(`/\*[\s\S]*?\*/`), skipHandler},

			// Keywords (must come before identifiers)
			{regexp.MustCompile(`\bdisplay\b`), defaultHandler(DISPLAY)},
			{regexp.MustCompile(`\bpixel\b`), defaultHandler(PIXEL)},
			{regexp.MustCompile(`\bprint\b`), defaultHandler(PRINT)},
			{regexp.MustCompile(`\bpush\b`), defaultHandler(PUSH)},
			{regexp.MustCompile(`\brgba\b`), defaultHandler(RGBA_FUNC)},
			{regexp.MustCompile(`\bhex\b`), defaultHandler(HEX_FUNC)},

			// Literals
			{regexp.MustCompile(`\d+`), defaultHandler(NUMBER)},
			{regexp.MustCompile(`"[^"]*"`), defaultHandler(STRING)},
			{regexp.MustCompile(`#[0-9a-fA-F]{6}`), defaultHandler(HEX_LITERAL)},
			{regexp.MustCompile(`\[(\s*\d+\s*,?){4}\]`), defaultHandler(RGBA_LITERAL)},

			// Identifiers
			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`), defaultHandler(IDENTIFIER)},

			// Symbols (ordered by length to avoid conflicts)
			{regexp.MustCompile(`\.`), defaultHandler(DOT)},
			{regexp.MustCompile(`:`), defaultHandler(COLON)},
			{regexp.MustCompile(`,`), defaultHandler(COMMA)},
			{regexp.MustCompile(`=`), defaultHandler(ASSIGN)},
			{regexp.MustCompile(`\(`), defaultHandler(LPAREN)},
			{regexp.MustCompile(`\)`), defaultHandler(RPAREN)},
			{regexp.MustCompile(`\[`), defaultHandler(LBRACKET)},
			{regexp.MustCompile(`\]`), defaultHandler(RBRACKET)},
		},
	}
}
