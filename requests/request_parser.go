package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/url"
	"strings"

	"github.com/simple-url/surl/utils"
)

func JsonRequestParser(payload map[string]interface{}) (io.Reader, string, error) {
	json_byte, err := json.Marshal(payload)
	if err != nil {
		return nil, "", err
	}
	json_reader := bytes.NewReader(json_byte)
	return json_reader, "application/json", nil
}

type FormUrlItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func FormUrlRequestParser(payload []FormUrlItem) (io.Reader, string, error) {
	form := url.Values{}
	for _, item := range payload {
		form.Set(item.Name, item.Value)
	}
	return bytes.NewBufferString(form.Encode()), "application/x-www-form-urlencoded", nil
}

type FormMultipartItem struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Value    *string `json:"value"`
	FileName *string `json:"file_name"`
	FilePath *string `json:"file_path"`
}

func FormMultipartRequestParser(payload []FormMultipartItem, file_reader utils.IFileReader) (io.Reader, string, error) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	defer writer.Close()
	for _, item := range payload {
		var fw io.Writer
		var err error
		if item.Type == "string" {
			if item.Value == nil {
				return nil, "", errors.New("if type = 'string', value must be defined")
			}
			if fw, err = writer.CreateFormField(item.Name); err != nil {
				return nil, "", err
			}
			if _, err := io.Copy(fw, strings.NewReader(*item.Value)); err != nil {
				return nil, "", err
			}
		} else if item.Type == "file" {
			if item.FilePath == nil || item.FileName == nil {
				return nil, "", errors.New("if type = 'file', file_path and file_name must be defined")
			}
			if fw, err = writer.CreateFormFile(item.Name, *item.FileName); err != nil {
				return nil, "", err
			}
			file_data, err := file_reader.ReadFile(*item.FilePath)
			if err != nil {
				return nil, "", err
			}
			if file_data == nil {
				return nil, "", errors.New("file " + *item.FilePath + " not found")
			}
			if _, err := io.Copy(fw, bytes.NewReader(*file_data)); err != nil {
				return nil, "", err
			}
		} else {
			return nil, "", errors.New("unknown type: " + item.Type + ", type value must be string/file")
		}
	}
	return buf, writer.FormDataContentType(), nil
}
