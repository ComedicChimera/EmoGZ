package interpreter

import (
	"github.com/EmoGZ/scanner"
	"fmt"
	"os"
	"strconv"
	"log"
)

var commands = map[string]func(args []scanner.Token) scanner.Token{
	"PRINT": printOperand,
	"SET": setOperand,
	"GET": getOperand,
	"SCAN": scanInput,
	"EXIT": stop,
	"PUSH": pushToCache,
	"POP": popFromCache,
	"INVERT": changeSine,
	"INCREMENT": incrementOperand,
	"DECREMENT": decrementOperand,
	"EQUAL": compareOperand,
}

var operand = scanner.Token{"INTEGER", "0", -1}
var cache []scanner.Token

var null = scanner.Token{}

func Run(command string, args []scanner.Token) scanner.Token {
	// fmt.Println(command, operand)
	return commands[command](args)
}

func printOperand(args []scanner.Token) scanner.Token {
	if len(args) > 0 {
		log.Fatal("😂 takes no parameters.")
	}
	fmt.Printf(operand.Value)
	return null
}

func setOperand(args []scanner.Token) scanner.Token {
	if len(args) != 1 {
		log.Fatal("💯 takes 1 parameter.")
	}
	operand = args[0]
	return null
}

func getOperand(args []scanner.Token) scanner.Token {
	if len(args) != 0 {
		log.Fatal("👀 takes no parameters.")
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
		log.Fatal("Failed to read input from the STDIN.")
	}
	return scanner.Token{Name: "CHAR", Value: string(inp), Index: -1}
}

func stop(args []scanner.Token) scanner.Token {
	if len(args) != 0 {
		log.Fatal("💀 takes no parameters.")
	}
	os.Exit(0)
	return null
}

func pushToCache(args []scanner.Token) scanner.Token {
	cache = append(args, cache...)
	if len(cache) > 3000 {
		log.Fatal("Cache overflow.")
	}
	return null
}

func popFromCache(args []scanner.Token) scanner.Token {
	if len(args) != 1 {
		log.Fatal("🍦 takes 1 parameter.")
	}
	if args[0].Name != "INTEGER" {
		log.Fatal("Cache index must be an integer.")
	}
	val, _ := strconv.Atoi(args[0].Value)
	if val < 0 || val >= len(cache) {
		log.Fatal("Invalid cache index.")
	}
	operand = cache[val]
	cache = append(cache[:val], cache[val+1:]...)
	return null
}

func changeSine(args []scanner.Token) scanner.Token {
	if len(args) != 0 {
		log.Fatal("🔫 takes no parameters.")
	}
	if operand.Name == "INTEGER" {
		val, _ := strconv.Atoi(operand.Value)
		operand.Value = strconv.Itoa(-val)
	} else {
		val := []rune(operand.Value)[0]
		operand.Value = strconv.Itoa(-int(val))
		operand.Name = "INTEGER"
	}
	return null
}

func incrementOperand(args []scanner.Token) scanner.Token {
	if len(args) != 0 {
		log.Fatal("😺 takes no parameters.")
	}
	if operand.Name == "INTEGER" {
		val, _ := strconv.Atoi(operand.Value)
		operand.Value = strconv.Itoa(val + 1)
	} else {
		val := []rune(operand.Value)[0]
		operand.Value = strconv.Itoa(int(val) + 1)
		operand.Name = "INTEGER"
	}
	return null
}

func decrementOperand(args []scanner.Token) scanner.Token {
	if len(args) != 0 {
		log.Fatal("💩 takes no parameters.")
	}
	if operand.Name == "INTEGER" {
		val, _ := strconv.Atoi(operand.Value)
		operand.Value = strconv.Itoa(val - 1)
	} else {
		val := []rune(operand.Value)[0]
		operand.Value = strconv.Itoa(int(val) - 1)
		operand.Name = "INTEGER"
	}
	return null
}

func compareOperand(args []scanner.Token) scanner.Token {
	if len(args) < 1 {
		log.Fatal("💙 takes at least 1 parameter.")
	}
	for _, token := range args {
		if token.Value != operand.Value || token.Name != operand.Name {
			return scanner.Token{Name:"INTEGER", Value:"0", Index: -1}
		}
	}
	return scanner.Token{Name:"INTEGER", Value:"1", Index: -1}
}

func IfOperand() bool {
	if operand.Name == "INTEGER" {
		val, _ := strconv.Atoi(operand.Value)
		if val > 0 {
			return true
		}
	} else {
		return true
	}
	return false
}