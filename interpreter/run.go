package interpreter

import (
	"github.com/EmoGZ/scanner"
	"fmt"
	"os"
)

var commands = map[string]func(args []scanner.Token) scanner.Token{
	"PRINT": printOperand,
	"SET": setOperand,
	"GET": getOperand,
	"SCAN": scanInput,
}

var operand = scanner.Token{"INTEGER", "0", -1}

func Run(command string, args []scanner.Token) scanner.Token {
	return commands[command](args)
}

func printOperand(args []scanner.Token) scanner.Token {
	if len(args) > 0 {
		fmt.Println("😂 takes no parameters.")
		os.Exit(1)
	}
	if operand.Name == "CHAR" {
		fmt.Printf("'%s'\n", operand.Value)
	} else {
		fmt.Println(operand.Value)
	}
	return scanner.Token{}
}

func setOperand(args []scanner.Token) scanner.Token {
	if len(args) != 1 {
		fmt.Println("💯 takes 1 parameter.")
		os.Exit(1)
	}
	operand = args[0]
	return scanner.Token{}
}

func getOperand(args []scanner.Token) scanner.Token {
	if len(args) != 0 {
		fmt.Println("👀 takes no parameters.")
		os.Exit(1)
	}
	return operand
}

func scanInput(args []scanner.Token) scanner.Token {
	if len(args) != 0 {
		fmt.Println("👑 takes no parameters.")
		os.Exit(1)
	}
	var inp rune
	_, err := fmt.Scanf("%c", &inp)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to read input from the STDIN.")
		os.Exit(1)
	}
	return scanner.Token{Name: "CHAR", Value: string(inp), Index: -1}
}
