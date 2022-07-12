package httpHelper

import (
	"encoding/json"
	"net/http"
)

func JsonCodeResponse(response http.ResponseWriter, statusCode int) {
	response.WriteHeader(statusCode)
}

func JsonResponse(response http.ResponseWriter, responseModel interface{}) error {
	dataBytes, err := json.Marshal(&responseModel)
	if err != nil {
		return err
	}

	_, err = response.Write(dataBytes)
	if err != nil {
		return err
	}

	response.WriteHeader(http.StatusOK)

	return nil
}
