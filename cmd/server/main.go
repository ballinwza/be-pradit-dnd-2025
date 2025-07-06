package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/ballinwza/be-pradit-dnd-2025/cmd/cli"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/config"
	"github.com/ballinwza/be-pradit-dnd-2025/middleware"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	config.LoadingEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srvInjecter := cli.InjecterAllService()
	authHandler := middleware.AuthMiddleWare(srvInjecter)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FE_URL_DOMAIN")},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", corsHandler.Handler(authHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
