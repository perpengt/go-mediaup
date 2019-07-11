package mediaup

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"net/http"

	"github.com/perpengt/ids"
)

var (
	ErrNotPng = errors.New("media server only supports png file")
)

func UploadImage(url string, img image.Image) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	err := png.Encode(buf, img)
	if err != nil {
		return nil, err
	}

	return UploadImageBytes(url, buf.Bytes())
}

func UploadImageBytes(url string, data []byte) ([]byte, error) {
	if http.DetectContentType(data) != "image/png" {
		return nil, ErrNotPng
	}

	req, err := newUploadRequest(url, data)
	if err != nil {
		return nil, err
	}

	res, err := sendRequest(req)
	if err != nil {
		return nil, err
	}

	if !res.Ok {
		if len(res.ErrorMessage) > 0 {
			return nil, errors.New(res.ErrorMessage)
		} else {
			return nil, errors.New(res.ErrorCode)
		}
	}

	id, err := ids.DecodeID(res.Data)
	if err != nil {
		return nil, err
	}

	return id.Bytes(), nil
}
