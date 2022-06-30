package httpHelper

import "net/http"

func JsonResponse(response http.ResponseWriter, statusCode int) error {
	response.WriteHeader(statusCode)
	return nil
}
