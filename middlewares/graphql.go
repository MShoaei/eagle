package middlewares

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const (
	production = true
)

var (
	signKeyByte, _   = ioutil.ReadFile(`keys/private_key.pem`)
	verifyKeyByte, _ = ioutil.ReadFile(`keys/public_key.pem`)

	SignKey, _   = jwt.ParseRSAPrivateKeyFromPEM(signKeyByte)
	VerifyKey, _ = jwt.ParseRSAPublicKeyFromPEM(verifyKeyByte)
)

type AuthToken string

func GetAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := AuthToken("Authorization")
		ctx := context.WithValue(r.Context(), authorization, r.Header.Get("Authorization"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
