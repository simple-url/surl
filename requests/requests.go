package requests

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"
)

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
