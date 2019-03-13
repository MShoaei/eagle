package middlewares

import (
	"net/http"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	r.Header.Get("Authorization")
}
