package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph"

	character_service "github.com/ballinwza/be-pradit-dnd-2025/internal/outbound/character/service"
	user_service "github.com/ballinwza/be-pradit-dnd-2025/internal/outbound/user/service"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/config"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	config.LoadingEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userService := user_service.NewUserService()
	characterService := character_service.NewCharacterService()
	
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserService: userService,
		CharacterService: characterService,
	}}))

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FE_URL_DOMAIN")},
        AllowCredentials: true,
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"*"},
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", corsHandler.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
