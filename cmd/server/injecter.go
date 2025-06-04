package main

import (
	"github.com/99designs/gqlgen/graphql/handler"

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

	return srv
}