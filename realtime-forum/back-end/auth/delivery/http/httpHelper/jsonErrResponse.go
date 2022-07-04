package httpHelper

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func ErrJsonResponse(response http.ResponseWriter, msg string, code int) error {
	errModel := &ErrResponse{
		Msg:  msg,
		Code: code,
	}

	log.Println("errr 123")
	dataBytes, err := json.Marshal(errModel)
	if err != nil {
		return err
	}

	_, err = response.Write(dataBytes)
	if err != nil {
		return err
	}

	response.WriteHeader(code)

	return nil
}
