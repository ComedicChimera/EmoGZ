package scanner

import (
	"strconv"
	"log"
)

type Token struct {
	Name string
	Value string
	Index int
}

func Scan(text string) []Token {
	tokenMap := map[string]string {
		"😂": "PRINT",
		"🔫": "INVERT",
		"💀": "EXIT",
		"💩": "DECREMENT",
		"😺": "INCREMENT",
		"👌": "COMPARE",
		"💙": "EQUAL",
		"👀": "GET",
		"💯": "SET",
		"🍦": "POP",
		"💎": "PUSH",
		"🍊": "LOOP",
		"🍌": "BREAK",
		"🏁": "END",
		"👑": "SCAN",
	}

	var tokens []Token
	var isChar bool
	var isInt bool
	var currentToken string
	var char string

	for i, j := range text {
		char = string(j)
		if isChar || isInt {
			if char == "'" {
				tokens = append(tokens, Token{"CHAR", currentToken, i})
				isChar = false
				currentToken = ""
			} else if _, err := strconv.Atoi(char); err != nil && isInt {
				if _, err := strconv.Atoi(currentToken); err != nil {
					log.Fatalf("Invalid integer. (%s at position %d)", char, i)
				}
				tokens = append(tokens, Token{"INTEGER", currentToken, i})
				isInt = false
				currentToken = ""
			} else {
				currentToken += char
			}
		} else if val, ok := tokenMap[char]; ok {
			tokens = append(tokens, Token{val, char, i})
		} else if char == "'" {
			isChar = true
		} else if _, err := strconv.Atoi(char); err == nil {
			currentToken += char
			isInt = true
		}
	}

	for _, token := range tokens {
		if token.Name == "CHAR" {
			if len([]rune(token.Value)) != 1 {
				log.Fatalf("Invalid character '%s' at position %d.", token.Value, token.Index)
			}
		}
	}

	return tokens
}
