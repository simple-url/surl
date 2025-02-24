package utils

import (
	"fmt"
	"strings"
)

func PrintWithWhiteSpace(text string, max_space int) string {
	var sb strings.Builder
	text_length := len(text)
	for i := 0; i < max_space; i++ {
		if i < text_length {
			sb.WriteString(string(text[i]))
		} else {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func HelpMessage() {
	fmt.Println("SURL v0.0.1")
	fmt.Println()
	fmt.Println("Commands:")
	// fmt.Println("i init       create surl.json")
	fmt.Println("  list       show list of requests")
	fmt.Println("  run <name> run http request by name")
	fmt.Println("  help       show this help message")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -h         show help for specific command")
}
