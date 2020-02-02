package exception

import "net/http"

type Exception struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

var NotFound = Exception{
	Message:    "Not found",
	StatusCode: http.StatusNotFound,
}

var BadRequest = Exception{
	Message:    "invalid request",
	StatusCode: http.StatusBadRequest,
}
