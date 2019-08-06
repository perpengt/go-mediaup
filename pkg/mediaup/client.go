package mediaup

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"github.com/perpengt/mediaup/internal/resp"
)

func newUploadRequest(url string, data []byte) (*http.Request, error) {
	var (
		mimeType string
		fileExt string
	)

	// Detect mime type
	mimeType = http.DetectContentType(data)
	switch mimeType {
	case "image/png":
		fileExt = ".png"

	case "image/jpeg":
		fileExt = ".jpg"

	default:
		return nil, ErrUnsupportedFileType
	}

	// Create buffer
	buf := bytes.NewBuffer([]byte{})
	mw := multipart.NewWriter(buf)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", "form-data; name=\"file\"; filename=\"upload" + fileExt + "\"")
	h.Set("Content-Type", mimeType)

	part, err := mw.CreatePart(h)
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
