package requests_test

import (
	"io"
	"regexp"
	"strings"
	"testing"

	"github.com/simple-url/surl/requests"
)

func TestJsonParser(t *testing.T) {
	json_payload := map[string]interface{}{
		"string": "string",
		"int":    10,
		"bool":   true,
		"array":  []string{"a", "b", "c"},
		"object": map[string]interface{}{
			"foo":   "bar",
			"hello": 20,
		},
	}
	data, ct, err := requests.JsonRequestParser(json_payload)
	if err != nil {
		t.Error(err.Error())
	}
	data_string := new(strings.Builder)
	_, err = io.Copy(data_string, data)
	if err != nil {
		t.Error(err.Error())
	}
	if data_string.String() != "{\"array\":[\"a\",\"b\",\"c\"],\"bool\":true,\"int\":10,\"object\":{\"foo\":\"bar\",\"hello\":20},\"string\":\"string\"}" {
		t.Error("Invalid json parser output")
	}
	if ct != "application/json" {
		t.Error("wrong content type")
	}
}

func TestFormUrlParser(t *testing.T) {
	data, ct, err := requests.FormUrlRequestParser([]requests.FormUrlItem{
		{
			Name:  "username",
			Value: "testuser",
		},
		{
			Name:  "password",
			Value: "somepassword",
		},
	})
	if err != nil {
		t.Error(err)
	}

	data_string := new(strings.Builder)
	_, err = io.Copy(data_string, data)
	if err != nil {
		t.Error(err.Error())
	}
	if data_string.String() != "password=somepassword&username=testuser" {
		t.Error("Invalid form url parser output")
	}
	if ct != "application/x-www-form-urlencoded" {
		t.Error("wrong content type")
	}
}

type MockFileReader struct {
}

func (fr *MockFileReader) ReadFile(file_path string) (*[]byte, error) {
	res := []byte("this example of file content")
	return &res, nil
}

func TestFormMultipartParser(t *testing.T) {
	value := "testuser"
	file_name := "README.md"
	file_path := "./README.md"
	data, ct, err := requests.FormMultipartRequestParser([]requests.FormMultipartItem{
		{
			Name:  "username",
			Type:  "string",
			Value: &value,
		},
		{
			Name:     "my_file",
			Type:     "file",
			FileName: &file_name,
			FilePath: &file_path,
		},
	}, &MockFileReader{})
	if err != nil {
		t.Error(err)
	}

	data_string := new(strings.Builder)
	_, err = io.Copy(data_string, data)
	if err != nil {
		t.Error(err.Error())
	}
	match, err := regexp.MatchString("--(?s:.)+Content-Disposition: form-data; name=\"username\"(?s:.)*testuser(?s:.)*--", data_string.String())
	if err != nil {
		t.Error(err)
	}
	if !match {
		t.Error("username no set")
	}
	match, err = regexp.MatchString("--(?s:.)+Content-Disposition: form-data; name=\"my_file\"(?s:.)*this example of file content(?s:.)*--", data_string.String())
	if err != nil {
		t.Error(err)
	}
	if !match {
		t.Error("my file no set")
	}

	if match, err = regexp.MatchString("multipart/form-data", ct); !match || err != nil {
		t.Error("wrong content type")
	}
}
