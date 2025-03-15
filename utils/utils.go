package utils

import (
	"errors"
	"fmt"
	"os"
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

type IResponse interface {
	Println(text string)
	Error(text string)
	Exit()
}

type Response struct {
}

func (r *Response) Println(text string) {
	fmt.Println(text)
}

func (r *Response) Error(text string) {
	fmt.Println(text)
	os.Exit(1)
}

func (r *Response) Exit() {
	os.Exit(1)
}

type MockResponse struct {
	Ok     []string
	Err    []string
	IsExit bool
}

func (r *MockResponse) Println(text string) {
	r.Ok = append(r.Ok, text)
}

func (r *MockResponse) Error(text string) {
	r.Err = append(r.Err, text)
	r.IsExit = true
}

func (r *MockResponse) Exit() {
	r.IsExit = true
}
