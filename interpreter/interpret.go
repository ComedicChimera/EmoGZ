package interpreter

import (
	"github.com/EmoGZ/scanner"
	"log"
)

func Execute(tokens []scanner.Token) bool {
	var operands []scanner.Token
	var awaitingCommands []string
	var pos int
	for pos < len(tokens) {
		token := tokens[pos]
		if token.Name == "END" {
			log.Fatalf("Unexpected 🏁 at %d", pos)
		} else if token.Name == "BREAK" {
			return true
		} else if token.Name == "LOOP" {
			tkRegion := selectRegion(pos, tokens)
			for !Execute(tkRegion) {}
			pos += len(tkRegion) + 1
		} else if token.Name == "INTEGER" || token.Name == "CHAR" {
			operands = append(operands, token)
		} else if len(operands) > 0 && len(awaitingCommands) > 0 {
			val := Run(awaitingCommands[len(awaitingCommands)-1], operands)
			if returns(awaitingCommands[len(awaitingCommands)-1]) {
				operands = []scanner.Token{val}
			} else {
				operands = []scanner.Token{}
			}
			awaitingCommands = awaitingCommands[:len(awaitingCommands)-1]
			continue
		} else if needsArgs(token.Name) {
			awaitingCommands = append(awaitingCommands, token.Name)
		} else if returns(token.Name) {
			val := Run(token.Name, operands)
			operands = []scanner.Token{val}
		} else if token.Name == "COMPARE" {
			region := selectRegion(pos, tokens)
			if IfOperand() {
				Execute(region)
			} else {
				pos += len(region) + 1
			}
 		} else {
			operands = []scanner.Token{}
			Run(token.Name, operands)
		}
		pos++
	}
	for len(awaitingCommands) > 0 {
		if len(operands) > 0 {
			val := Run(awaitingCommands[len(awaitingCommands)-1], operands)
			if returns(awaitingCommands[len(awaitingCommands)-1]) {
				operands = []scanner.Token{val}
			} else {
				operands = []scanner.Token{}
			}
			awaitingCommands = awaitingCommands[:len(awaitingCommands)-1]
		} else {
			log.Fatalf("Unsatisfied operator %s", awaitingCommands[0])
		}
	}
	return false
}

func returns(command string) bool {
	switch command {
	case "GET", "SCAN", "EQUAL":
		return true
	}
	return false
}

func needsArgs(command string) bool {
	switch command {
	case "PUSH", "POP", "EQUAL", "SET":
		return true
	}
	return false
}

func selectRegion(pos int, tokens []scanner.Token) []scanner.Token {
	lpos := len(tokens) - 1
	fpos := -1
	for lpos > pos {
		if tokens[lpos].Name == "END" {
			fpos = lpos
			break
		}
		lpos--
	}
	if fpos < 0 {
		log.Fatalf("%s region has no closer (at %d)", tokens[pos].Value, tokens[pos].Index)
	}
	return tokens[pos+1:fpos]
}