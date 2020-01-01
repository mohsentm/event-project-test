package responseMapper

import (
	"encoding/json"
	"net/http"
)

func ResponseHandler(w http.ResponseWriter, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(responseData)
}
