package responseMapper

import (
	"encoding/json"
	"net/http"
	"github.com/mohsentm/event-project-test/internal/helper/exception"
)

func Success(w http.ResponseWriter, responseData interface{}) {
	Json(w, responseData, http.StatusOK)
}

func Exception(w http.ResponseWriter, exception exception.Exception) {
	Json(w, exception, exception.StatusCode)
}

func Json(w http.ResponseWriter, responseData interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(responseData)
}
