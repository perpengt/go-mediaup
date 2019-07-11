package resp

import (
	"encoding/json"
	"io"
	"net/http"
)

type Data struct {
	Ok           bool   `json:"ok"`
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	Data         string `json:"data,omitempty"`
}

func Send(w http.ResponseWriter, status int, resp *Data) {
	if status > 100 {
		w.WriteHeader(status)
	}
	data, err := json.Marshal(resp)
	if err != nil {
		io.WriteString(w, `{"ok":false,"error_code":"server_error"}`)
	}
	w.Write(data)
}
