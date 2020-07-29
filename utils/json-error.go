package utils

import (
	"encoding/json"
	"net/http"
)

//HTTPErrorMessage type to represent custom HTTP error message
type HTTPErrorMessage struct {
	Error   string `json:"error"`
	Details string `json:"details"`
}

//JSONError creates HTTP response for error with JSON body
func JSONError(w http.ResponseWriter, err string, details string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(newHTTPErrorMessage(err, details))
}

func newHTTPErrorMessage(err string, details string) HTTPErrorMessage {
	return HTTPErrorMessage{
		Error:   err,
		Details: details,
	}
}
