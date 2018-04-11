package scanner

import (
	"regexp"
	"log"
)

func RemoveComments(text string) string {
	re, err := regexp.Compile("#[^\n]*\n")
	if err != nil {
		log.Fatal(err)
	}
	text = re.ReplaceAllStringFunc(text, replaceComment)
	return text
}

func replaceComment(comment string) string {
	rStr := ""
	for i := 0; i < len(comment) - 1; i++ {
		rStr += " "
	}
	return rStr + "\n"
}