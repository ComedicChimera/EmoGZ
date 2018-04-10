package main

import (
  	"fmt"
  	"os"
  	"golang.org/x/text/encoding/unicode"
  	"io"
  	"golang.org/x/text/transform"
	"io/ioutil"
	"github.com/EmoGZ/scanner"
	"github.com/EmoGZ/interpreter"
	"log"
)

func main() {
  if len(os.Args) > 1 {
    file, err := os.Open(os.Args[1])
    if err != nil {
    	log.Fatal("The specified file does not exist.")
	}

	r := NewUnicodeReader(file)

	utf8, _ := ioutil.ReadAll(r)
	text := string(utf8)
	tokens := scanner.Scan(text)
	interpreter.Execute(tokens)
  } else {
    fmt.Println("Please provide a valid file argument.")
  }
}

func NewUnicodeReader(file io.Reader) io.Reader {
	winutf := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	decoder := winutf.NewDecoder()
	return transform.NewReader(file, unicode.BOMOverride(decoder))
}
