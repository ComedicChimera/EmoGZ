package interpreter

import (
	"github.com/EmoGZ/scanner"
	"fmt"
	"os"
	"strconv"
)

var commands = map[string]func(args []scanner.Token) scanner.Token{
	"PRINT": printOperand,
	"SET": setOperand,
	"GET": getOperand,
	"SCAN": scanInput,
	"EXIT": stop,
	"PUSH": pushToCache,
	"POP": popFromCache,
}

var operand = scanner.Token{"INTEGER", "0", -1}
var cache []scanner.Token

var null = scanner.Token{}

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
	return null
}

func setOperand(args []scanner.Token) scanner.Token {
	if len(args) != 1 {
		fmt.Println("💯 takes 1 parameter.")
		os.Exit(1)
	}
	operand = args[0]
	return null
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

func stop(args []scanner.Token) scanner.Token {
	if len(args) != 0 {
		fmt.Println("💀 takes no parameters.")
		os.Exit(1)
	}
	os.Exit(0)
	return null
}

func pushToCache(args []scanner.Token) scanner.Token {
	cache = append(args, cache...)
	if len(cache) > 3000 {
		fmt.Println("Cache overflow.")
		os.Exit(1)
	}
	return null
}

func popFromCache(args []scanner.Token) scanner.Token {
	if len(args) != 1 {
		fmt.Println("🍦 takes 1 parameter.")
		os.Exit(1)
	}
	if args[0].Name != "INTEGER" {
		fmt.Println("Cache index must be an integer.")
		os.Exit(1)
	}
	val, _ := strconv.Atoi(args[0].Value)
	if val < 0 || val >= len(cache) {
		fmt.Println("Invalid cache index.")
		os.Exit(1)
	}
	operand = cache[val]
	cache = append(cache[:val], cache[val+1:]...)
	return null
}