package mediaup

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/perpengt/mediaup/internal/resp"
)

func newUploadRequest(url string, data []byte) (*http.Request, error) {
	buf := bytes.NewBuffer([]byte{})

	mw := multipart.NewWriter(buf)
	part, err := mw.CreateFormField("file")
	if err != nil {
		return nil, err
	}

	_, err = part.Write(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", mw.FormDataContentType())

	return req, nil
}

func sendRequest(req *http.Request) (*resp.Data, error) {
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data := new(resp.Data)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}
