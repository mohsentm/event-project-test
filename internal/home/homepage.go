package home

import (
	"fmt"
	"net/http"
)

func init() {

}
func Handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome home!")
}
