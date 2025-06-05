package cli

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"

	ability_detail_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/ability_detail/service"
	character_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/service"
	user_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/service"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph"
)

func InjecterAllService() *handler.Server{
	userService := user_service.NewUserService()
	characterService := character_service.NewCharacterService()
	abilityDetailService := ability_detail_service.NewAbilityDetailService()
	
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserService: userService,
		CharacterService: characterService,
		AbilityDetailService: abilityDetailService,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	srv.SetErrorPresenter(func (ctx context.Context, err error) *gqlerror.Error  {
		gqlErr := graphql.DefaultErrorPresenter(ctx, err)
		if gqlErr.Extensions == nil {
			gqlErr.Extensions = map[string]interface{}{}
		}

		gqlErr.Extensions["timestamp"] = time.Now().Format(time.RFC3339)

		if strings.Contains(gqlErr.Message, "does not exist in") {
			gqlErr.Extensions["friendly_message"] = "Error grqphql ใส่ field ไม่ถูกต้อง"
			gqlErr.Extensions["status_code"] = http.StatusBadRequest
		}

		return gqlErr
	})

	return srv
}