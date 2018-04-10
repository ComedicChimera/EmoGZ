package interpreter

import (
	"github.com/EmoGZ/scanner"
)

func Execute(tokens []scanner.Token) {
	var operands []scanner.Token
	var awaitingCommands []string
	var pos int
	for pos < len(tokens) {
		token := tokens[pos]
		if token.Name == "INTEGER" || token.Name == "CHAR"{
			operands = append(operands, token)
		} else if len(operands) > 0 && len(awaitingCommands) > 0{
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
		} else {
			Run(token.Name, operands)
		}
		pos++
	}
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
	case "PUSH", "POP", "EQUAL", "COMPARE", "SET":
		return true
	}
	return false
}