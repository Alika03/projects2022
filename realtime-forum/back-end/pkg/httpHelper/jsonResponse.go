package httpHelper

import "net/http"

func JsonCodeResponse(response http.ResponseWriter, statusCode int) {
	response.WriteHeader(statusCode)
}
