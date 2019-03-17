package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/MShoaei/eagle/api"
	"github.com/MShoaei/eagle/middlewares"
	"github.com/MShoaei/eagle/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/kataras/muxie"
)

const defaultPort = "9990"

func main() {
	mux := muxie.NewMux()
	cfg := api.Config{Resolvers: &api.Resolver{DB: models.DB}}

	cfg.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		tokenString := ctx.Value(middlewares.AuthToken("Authorization")).(string)
		tokenString, err = request.AuthorizationHeaderExtractor.Filter(tokenString)
		if err != nil {
			return nil, err
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("incorrect algorithm")
			}
			return middlewares.VerifyKey, nil
		})
		if err != nil {
			return nil, err
		}
		if !token.Valid {
			return nil, fmt.Errorf("invalid token")
		}
		return next(ctx)
	}

	mux.Use(middlewares.GetAuthMiddleware)
	mux.Handle("/", muxie.Methods().
		HandleFunc(http.MethodGet,
			handler.Playground("GraphQL playground", "/api")))
	mux.Handle("/api", muxie.Methods().
		Handle(http.MethodPost,
			handler.GraphQL(api.NewExecutableSchema(cfg))))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
