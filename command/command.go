package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/simple-url/surl/requests"
	"github.com/simple-url/surl/utils"
)

type SurlHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SurlRequest struct {
	Name    string       `json:"name"`
	Url     string       `json:"url"`
	Method  string       `json:"method"`
	Headers []SurlHeader `json:"headers"`
	Timeout *int         `json:"timeout"`
	Body    *string      `json:"body"`
}

type Surl struct {
	Requests []SurlRequest `json:"requests"`
}

func NewSurl() *Surl {
	return &Surl{
		Requests: []SurlRequest{},
	}
}

func (s *Surl) ReadJson(json_path string) error {
	file_byte, err := os.ReadFile(json_path)
	if err != nil {
		return err
	}
	s.Requests = []SurlRequest{}
	err = json.Unmarshal(file_byte, &s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Surl) List() {
	// Find max string length for println
	max_name_length := 4 + 1
	max_url_length := 3 + 1
	max_method_length := 6 + 1
	for _, item := range s.Requests {
		if len(item.Name)+1 > max_name_length {
			max_name_length = len(item.Name)
		}
		if len(item.Url)+1 > max_url_length {
			max_url_length = len(item.Url)
		}
		if len(item.Method)+1 > max_method_length {
			max_method_length = len(item.Method)
		}
	}
	// show data
	fmt.Println(utils.PrintWithWhiteSpace("NAME", max_name_length), utils.PrintWithWhiteSpace("METHOD", max_method_length), utils.PrintWithWhiteSpace("URL", max_url_length))
	for _, item := range s.Requests {
		fmt.Println(utils.PrintWithWhiteSpace(item.Name, max_name_length), utils.PrintWithWhiteSpace(item.Method, max_method_length), utils.PrintWithWhiteSpace(item.Url, max_url_length))
	}
}

func (s *Surl) ListHelp() {
	fmt.Println()
	fmt.Println("Usage: surl list <optional flags>")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -p <path>  overide json path (default: ./surl.json)")
	fmt.Println("  -h         show this help message")
}

func (s *Surl) Run(name string, verbose bool) error {
	is_found := false
	for _, item := range s.Requests {
		if item.Name == strings.TrimSpace(name) {
			// set mode
			is_found = true
			if verbose {
				fmt.Println(item.Name)
				fmt.Println(item.Url)
			}

			// set headers
			headers := map[string]string{}
			for _, x := range item.Headers {
				headers[x.Key] = x.Value
			}

			// set body
			var body io.Reader = nil
			if item.Body != nil {
				// if body != nil {
				// 	return errors.New(fmt.Sprint("invalid request on:", item.Name, "single request can only have one body/form/json"))
				// }
				body = strings.NewReader(*item.Body)
			}

			// make request
			var client = &http.Client{}
			var response *http.Response
			var err error
			if item.Timeout == nil {
				response, err = requests.MakeRequest(client, item.Method, item.Url, headers, body)
			} else {
				response, err = requests.MakeRequestWithTimeout(client, item.Method, item.Url, headers, body, *item.Timeout)
			}
			if err != nil {
				return err
			}
			defer response.Body.Close()

			res, err := requests.ResponseToString(response)
			if err != nil {
				return err
			}
			fmt.Println(res)
		}
	}
	if !is_found {
		return errors.New(fmt.Sprint("request with name ", name, " not found"))
	}
	return nil
}

func (s *Surl) RunHelp() {
	fmt.Println()
	fmt.Println("Usage: surl run <name> <optional flags>")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -p <path>  overide json path (default: ./surl.json)")
	fmt.Println("  -v         run verbosely")
	fmt.Println("  -h         show this help message")
}
