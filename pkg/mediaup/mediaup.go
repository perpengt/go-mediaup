package mediaup

import (
	"bytes"
	"errors"
	"github.com/perpengt/ids"
	"image"
	"image/png"
)

func UploadImage(url string, token string, img image.Image) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	err := png.Encode(buf, img)
	if err != nil {
		return nil, err
	}

	return UploadImageBytes(url, token, buf.Bytes())
}

func UploadImageBytes(url string, token string, data []byte) ([]byte, error) {

	req, err := newUploadRequest(url, token, data)
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
