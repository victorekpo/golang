package routes

import (
	"fmt"
	"net/http"
)

func HealthController(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "OK")
	if err != nil {
		return
	}
}
