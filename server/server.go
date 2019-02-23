package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MShoaei/command_control/models"

	"github.com/kataras/muxie"

	"github.com/99designs/gqlgen/handler"
	"github.com/MShoaei/command_control"
)

const defaultPort = "9990"

func main() {
	// router := gin.Default()
	router := muxie.NewMux()

	router.Handle("/", handler.Playground("GraphQL playground", "/api"))
	router.Handle("/api", handler.GraphQL(
		command_control.NewExecutableSchema(
			command_control.Config{
				Resolvers: &command_control.Resolver{
					DB: models.DB,
				}})))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
