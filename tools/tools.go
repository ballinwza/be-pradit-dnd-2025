//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/joho/godotenv"
	_ "github.com/rs/cors"
	_ "go.mongodb.org/mongo-driver/v2/mongo"
)