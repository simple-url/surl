package requests

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type Requests struct {
	Url     string
	Method  string
	Client  *http.Client
	Headers map[string]string
	Body    io.Reader
	Timeout int
}

type RequestsConfig struct {
	Url     string
	Method  string
	Headers map[string]string
	Client  *http.Client
	Body    io.Reader
	Timeout int
}

func InitRequests() *Requests {
	client := http.Client{}
	return &Requests{
		Url:     "/",
		Method:  "GET",
		Client:  &client,
		Headers: map[string]string{},
		Body:    nil,
		Timeout: 0,
	}
}

func InitRequestsWithConfig(config RequestsConfig) *Requests {
	if config.Url == "" {
		config.Url = "/"
	}
	if config.Method == "" {
		config.Method = "GET"
	}
	if config.Client == nil {
		client := http.Client{}
		config.Client = &client
	}
	return &Requests{
		Url:     config.Url,
		Method:  config.Method,
		Client:  config.Client,
		Headers: config.Headers,
		Body:    config.Body,
		Timeout: config.Timeout,
	}
}

func (r *Requests) SetUrl(url string) {
	r.Url = url
}

func (r *Requests) SetMethod(method string) {
	r.Method = method
}

func (r *Requests) SetBody(body io.Reader) {
	r.Body = body
}

func (r *Requests) OverrideClient(client *http.Client) {
	r.Client = client
}

func (r *Requests) AddHeader(key string, value string) {
	r.Headers[key] = value
}

func (r *Requests) RemoveHeader(key string) bool {
	if _, found := r.Headers[key]; !found {
		return false
	}
	delete(r.Headers, key)
	return true
}

func (r *Requests) Json(payload any) error {
	r.AddHeader("content-type", "application/json")
	json_byte, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	r.Body = bytes.NewReader(json_byte)
	return nil
}

func (r *Requests) Exec() (*http.Response, error) {
	var response *http.Response
	var err error
	if r.Timeout == 0 {
		response, err = MakeRequest(r.Client, r.Method, r.Url, r.Headers, r.Body)
	} else {
		response, err = MakeRequestWithTimeout(r.Client, r.Method, r.Url, r.Headers, r.Body, r.Timeout)
	}

	return response, err
}

// ====================================================
func MakeRequest(client *http.Client, method string, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		request.Header.Set(key, val)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func MakeRequestWithTimeout(client *http.Client, method string, url string, headers map[string]string, body io.Reader, timeout int) (*http.Response, error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	request, err := http.NewRequestWithContext(ctx, method, url, body)
	for key, val := range headers {
		request.Header.Set(key, val)
	}

	go func() {
		time.Sleep(time.Second * time.Duration(timeout))
		cancel()
	}()

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func ResponseToString(response *http.Response) (string, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, response.Body)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
