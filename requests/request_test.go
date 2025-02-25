package requests_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/simple-url/surl/requests"
)

type PingResponse struct {
	Path    string              `json:"path"`
	Method  string              `json:"method"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

type JsonRequest struct {
	Hello string `json:"hello"`
}

func TestMakeRequest(t *testing.T) {
	// Given
	res := PingResponse{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res.Path = r.URL.Path
		res.Method = r.Method
		res.Headers = r.Header
		body_buf := new(strings.Builder)
		if _, err := io.Copy(body_buf, r.Body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("content-type", "text/plain")
			w.Write([]byte(err.Error()))
			return
		}
		res.Body = body_buf.String()
		payload, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("content-type", "text/plain")
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Add("content-type", "application/json")
		w.Write(payload)
	}))
	defer server.Close()
	client := server.Client()

	// When
	json_body := JsonRequest{
		Hello: "world",
	}
	var json_buf bytes.Buffer
	err := json.NewEncoder(&json_buf).Encode(json_body)
	if err != nil {
		t.Error(err.Error())
	}
	response, err := requests.MakeRequest(client, "POST", server.URL+"/hai", map[string]string{
		"content-type": "application/json",
	}, &json_buf)

	// Expect
	if err != nil {
		t.Error(err.Error())
	}
	if res.Path != "/hai" {
		t.Error("wrong path, path should return /hai")
	}
	if res.Method != "POST" {
		t.Error("wrong method method should return POST")
	}
	vals, is_found := res.Headers["Content-Type"]
	if !is_found {
		t.Error("headers content-type not found")
	}
	is_value_found := false
	for _, item := range vals {
		if item == "application/json" {
			is_value_found = true
		}
	}
	if !is_value_found {
		t.Error("value content-type: application-json not found")
	}
	if res.Body != "{\"hello\":\"world\"}\n" {
		t.Error("invalid body value")
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err.Error())
	}
	response_json := PingResponse{}
	if err := json.Unmarshal(body, &response_json); err != nil {
		t.Error(err.Error())
	}
}

func TestMakeRequestWithTimeout(t *testing.T) {
	// Given
	res := PingResponse{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res.Path = r.URL.Path
		res.Method = r.Method
		res.Headers = r.Header
		body_buf := new(strings.Builder)
		if _, err := io.Copy(body_buf, r.Body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("content-type", "text/plain")
			w.Write([]byte(err.Error()))
			return
		}
		res.Body = body_buf.String()
		payload, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("content-type", "text/plain")
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Add("content-type", "application/json")
		w.Write(payload)
	}))
	defer server.Close()
	client := server.Client()

	// When
	json_body := JsonRequest{
		Hello: "world",
	}
	var json_buf bytes.Buffer
	err := json.NewEncoder(&json_buf).Encode(json_body)
	if err != nil {
		t.Error(err.Error())
	}
	response, err := requests.MakeRequestWithTimeout(client, "POST", server.URL+"/hai", map[string]string{
		"content-type": "application/json",
	}, &json_buf, 2)

	// Expect
	if err != nil {
		t.Error(err.Error())
	}
	if res.Path != "/hai" {
		t.Error("wrong path, path should return /hai")
	}
	if res.Method != "POST" {
		t.Error("wrong method method should return POST")
	}
	vals, is_found := res.Headers["Content-Type"]
	if !is_found {
		t.Error("headers content-type not found")
	}
	is_value_found := false
	for _, item := range vals {
		if item == "application/json" {
			is_value_found = true
		}
	}
	if !is_value_found {
		t.Error("value content-type: application-json not found")
	}
	if res.Body != "{\"hello\":\"world\"}\n" {
		t.Error("invalid body value")
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err.Error())
	}
	response_json := PingResponse{}
	if err := json.Unmarshal(body, &response_json); err != nil {
		t.Error(err.Error())
	}
}

func TestMakeRequestWithTimeoutTimeout(t *testing.T) {
	// Given
	res := PingResponse{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res.Path = r.URL.Path
		res.Method = r.Method
		res.Headers = r.Header
		body_buf := new(strings.Builder)
		if _, err := io.Copy(body_buf, r.Body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("content-type", "text/plain")
			w.Write([]byte(err.Error()))
			return
		}
		res.Body = body_buf.String()
		payload, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("content-type", "text/plain")
			w.Write([]byte(err.Error()))
			return
		}
		time.Sleep(time.Second * time.Duration(3))
		w.Header().Add("content-type", "application/json")
		w.Write(payload)
	}))
	defer server.Close()
	client := server.Client()

	// When
	json_body := JsonRequest{
		Hello: "world",
	}
	var json_buf bytes.Buffer
	err := json.NewEncoder(&json_buf).Encode(json_body)
	if err != nil {
		t.Error(err.Error())
	}
	_, err = requests.MakeRequestWithTimeout(client, "POST", server.URL+"/hai", map[string]string{
		"content-type": "application/json",
	}, &json_buf, 2)

	// Expect
	if err == nil {
		t.Error("should return timeout error")
	}
}
