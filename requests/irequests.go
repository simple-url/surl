package requests

import (
	"io"
	"net/http"
)

type IRequests interface {
	MakeRequest(client *http.Client, method string, url string, headers map[string]string, body io.Reader) (*http.Response, error)
	MakeRequestWithTimeout(client *http.Client, method string, url string, headers map[string]string, body io.Reader, timeout int) (*http.Response, error)
	ResponseToString(response *http.Response) (string, error)
}
