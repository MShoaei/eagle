package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/MShoaei/command_control"
	_ "github.com/go-chi/chi"
)

const defaultPort = "9990"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/api"))
	http.Handle("/api", handler.GraphQL(command_control.NewExecutableSchema(command_control.Config{Resolvers: &command_control.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
