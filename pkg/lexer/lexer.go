package lexer

import (
	"bytes"

	"github.com/latiif/lail/pkg/token"
)

// Lexer parses input into token
type Lexer struct {
	input   string
	pos     int  // current pos in input
	readPos int  // current reading pos after current char
	ch      byte // char to examin
	line    int  // current line
	col     int  // current col
}

// New instantiates a new LExer
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, col: 1}
	l.readChar()
	return l
}

// NextToken returns next token for the Lexer's input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch), Line: l.line, Col: l.col - 1}
		} else {
			tok = newChToken(token.Assign, l.ch, l.line, l.col)
		}
	case ';':
		tok = newChToken(token.Semicolon, l.ch, l.line, l.col)
	case '(':
		tok = newChToken(token.Lparen, l.ch, l.line, l.col)
	case ')':
		tok = newChToken(token.Rparen, l.ch, l.line, l.col)
	case '{':
		tok = newChToken(token.Lbrace, l.ch, l.line, l.col)
	case '}':
		tok = newChToken(token.Rbrace, l.ch, l.line, l.col)
	case '[':
		tok = newChToken(token.Lbracket, l.ch, l.line, l.col)
	case ']':
		tok = newChToken(token.Rbracket, l.ch, l.line, l.col)
	case '+':
		tok = newChToken(token.Plus, l.ch, l.line, l.col)
	case '-':
		tok = newChToken(token.Minus, l.ch, l.line, l.col)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NEQ, Literal: string(ch) + string(l.ch), Line: l.line, Col: l.col - 1}
		} else {
			tok = newChToken(token.Bang, l.ch, l.line, l.col)
		}
	case '/':
		tok = newChToken(token.Slash, l.ch, l.line, l.col)
	case '*':
		tok = newChToken(token.Astersik, l.ch, l.line, l.col)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LTE, Literal: string(ch) + string(l.ch), Line: l.line, Col: l.col - 1}
		} else {
			tok = newChToken(token.LT, l.ch, l.line, l.col)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GTE, Literal: string(ch) + string(l.ch), Line: l.line, Col: l.col - 1}
		} else {
			tok = newChToken(token.GT, l.ch, l.line, l.col)
		}
	case ',':
		tok = newChToken(token.Comma, l.ch, l.line, l.col)
	case '.':
		tok = newChToken(token.Dot, l.ch, l.line, l.col)
	case '"':
		tok.Type = token.String
		tok.Line = l.line
		tok.Col = l.col
		tok.Literal = l.readString()
		return tok
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = l.line
		tok.Col = l.col
	default:
		if isLetter(l.ch) || l.ch == '_' {
			tok.Line = l.line
			tok.Col = l.col
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Line = l.line
			tok.Col = l.col
			tok.Type = token.Int
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newChToken(token.Illegal, l.ch, l.line, l.col)
	}
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
	if l.ch == '\n' {
		l.line++
		l.col = 0
	} else {
		l.col++
	}
}

func newChToken(tokenType token.Type, ch byte, line, col int) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
		Line:    line,
		Col:     col,
	}
}

func isLetter(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z')
}

func isDigit(ch byte) bool {
	return (ch >= '0' && ch <= '9')
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos
	if isLetter(l.ch) || l.ch == '_' {
		for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
			l.readChar()
		}
	}
	return l.input[pos:l.pos]
}

var escapeCharacters = map[byte]byte{
	'n':  '\n',
	'r':  '\r',
	'\\': '\\',
	'"':  '"',
	't':  '\t',
}

func (l *Lexer) readString() string {
	var str bytes.Buffer
	l.readChar()
	for l.ch != '"' {
		if l.ch == '\\' {
			l.readChar()
			if val, ok := escapeCharacters[l.ch]; ok {
				str.WriteByte(val)
			}
		} else {
			str.WriteByte(l.ch)
		}
		l.readChar()
	}
	l.readChar()
	return str.String()
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}
