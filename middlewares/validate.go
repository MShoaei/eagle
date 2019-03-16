package middlewares

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
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

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	//validate token
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("incorrect algorithm")
			}
			return VerifyKey, nil
		},
	)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorised access to this resource")
		return
	}

	if token.Valid {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token is not valid")
	}
}
