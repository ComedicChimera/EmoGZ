package interpreter

import (
	"github.com/EmoGZ/scanner"
	"fmt"
	"os"
)

func Execute(tokens []scanner.Token) {
	var operatorStack []string
	var operandStack []scanner.Token
	var awaitingOperand bool

	for _, token := range tokens {
		if token.Name == "INTEGER" || token.Name == "CHAR"{
			if awaitingOperand {
				operandStack = append(operandStack, token)
			} else {
				fmt.Printf("Unexpected Token [%s at %d]\n", token.Value, token.Index)
				os.Exit(1)
			}
		} else {
			if awaitingOperand {
				if returns(token.Name) {
					if needsArgs(token.Name) {
						operatorStack = append(operatorStack, token.Name)
					} else {
						awaitingOperand = false
						operandStack = append(operandStack, Run(token.Name, []scanner.Token{}))
					}
				} else {
					if len(operandStack) > 0 {
						if !needsArgs(token.Name) {
							awaitingOperand = false
						}
						Run(operatorStack[len(operatorStack) - 1], operandStack)
						operatorStack[len(operatorStack) - 1] = token.Name
						operandStack = []scanner.Token{}
					} else {
						fmt.Printf("Operator does not return a value [%s at %d]\n", token.Name, token.Index)
					}
				}
			} else {
				if len(operatorStack) > 0 {
					val := Run(operatorStack[len(operatorStack) - 1], operandStack)
					if returns(operatorStack[len(operatorStack) - 1]) {
						operandStack = append(operandStack, val)
					}
					operatorStack[len(operatorStack) - 1] = token.Name
				} else {
					operatorStack = append(operatorStack, token.Name)
					if needsArgs(token.Name) {
						awaitingOperand = true
					}
				}
			}
		}
	}
	if len(operatorStack) > 0 {
		Run(operatorStack[0], operandStack)
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