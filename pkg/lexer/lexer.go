package lexer

import (
	"github.com/latiif/lail/pkg/token"
)

// Lexer parses input into token
type Lexer struct {
	input   string
	pos     int  // current pos in input
	readPos int  // current reading pos after current char
	ch      byte // char to examin
}

// New instantiates a new LExer
func New(input string) *Lexer {
	l := &Lexer{input: input}
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
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newChToken(token.Assign, l.ch)
		}
	case ';':
		tok = newChToken(token.Semicolon, l.ch)
	case '(':
		tok = newChToken(token.Lparen, l.ch)
	case ')':
		tok = newChToken(token.Rparen, l.ch)
	case '{':
		tok = newChToken(token.Lbrace, l.ch)
	case '}':
		tok = newChToken(token.Rbrace, l.ch)
	case '+':
		tok = newChToken(token.Plus, l.ch)
	case '-':
		tok = newChToken(token.Minus, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newChToken(token.Bang, l.ch)
		}
	case '/':
		tok = newChToken(token.Slash, l.ch)
	case '*':
		tok = newChToken(token.Astersik, l.ch)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LTE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newChToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GTE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newChToken(token.GT, l.ch)
		}
	case ',':
		tok = newChToken(token.Comma, l.ch)
	case '"':
		tok.Type = token.String
		tok.Literal = l.readString()
		return tok
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.Int
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newChToken(token.Illegal, l.ch)
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
}

func newChToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z')
}

func isDigit(ch byte) bool {
	return (ch >= '0' && ch <= '9')
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readString() string {
	l.readChar()
	pos := l.pos
	for l.ch != '"' {
		l.readChar()
	}
	l.readChar()
	return l.input[pos : l.pos-1]
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
