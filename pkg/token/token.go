package token

// Type describe the token's type
type Type string

// Token represents a single token
type Token struct {
	Type    Type
	Literal string
	Line    int
	Col     int
}

var keywords = map[string]Type{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"import": Import,
	"return": Return,
}

const (
	// Illegal means an unknown token
	Illegal = "ILLEGAL"
	// EOF marks end of file
	EOF = "EOF"
	// Ident is for identifiers
	Ident = "IDENT" // add, foobar, x, y, ...
	// Int is Integer, eg 21312, -235
	Int = "INT"
	// Assign operator
	Assign = "="
	// Plus operator
	Plus = "+"
	// Minus operator
	Minus = "-"
	// Bang operator
	Bang = "!"
	// Astersik operator
	Astersik = "*"
	// Slash token
	Slash = "/"
	// Comma ,
	Comma = ","
	// Dot .
	Dot = "."
	// Semicolon ;
	Semicolon = ";"
	// Lparen is left parenthesis (
	Lparen = "("
	// Rparen is right parenthesis )
	Rparen = ")"
	// Lbrace is left brace {
	Lbrace = "{"
	// Rbrace is right brace }
	Rbrace = "}"
	// Lbracket is left bracket
	Lbracket = "["
	// Rbracket is right bracket
	Rbracket = "]"
	// LT (less than)
	LT = "<"
	// GT (greater than)
	GT = ">"
	// LTE (less than or equal)
	LTE = "<="
	// GTE (greater than or equal)
	GTE = ">="
	// EQ logical equality
	EQ = "=="
	// NEQ Not equal
	NEQ = "!="
	// Function declares a function
	Function = "FUNCTION"
	// Let declares a variable
	Let = "LET"
	// True boolean literal
	True = "TRUE"
	// False boolean literal
	False = "FALSE"
	// If conditional
	If = "IF"
	// Else else-conditional
	Else = "ELSE"
	// Return keyword
	Return = "RETURN"
	// DQuote is Double quotation
	DQuote = "\""
	// String
	String = "STRING"
	// Import keyword
	Import = "IMPORT"
)

// LookupIdent looks up a string in keywords
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Ident
}
