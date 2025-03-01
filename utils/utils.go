package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var VERSION string = "v1.0.0"

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
	fmt.Println("SURL " + VERSION)
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

type IFileReader interface {
	ReadFile(path string) (*[]byte, error)
}

type FileReader struct {
}

func (fr *FileReader) ReadFile(file_path string) (*[]byte, error) {
	if _, err := os.Stat(file_path); errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	file_data, err := os.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	return &file_data, nil
}
