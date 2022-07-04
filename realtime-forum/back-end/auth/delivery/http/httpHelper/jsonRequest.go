package httpHelper

import (
	"encoding/json"
	"io"
	"net/http"
)

func BindJson(request *http.Request, dto interface{}) error {
	dataBytes, err := io.ReadAll(request.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(dataBytes, dto)
}
