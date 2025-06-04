package main

import (
	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph"
	ability_detail_service "github.com/ballinwza/be-pradit-dnd-2025/internal/outbound/ability_detail/service"
	character_service "github.com/ballinwza/be-pradit-dnd-2025/internal/outbound/character/service"
	user_service "github.com/ballinwza/be-pradit-dnd-2025/internal/outbound/user/service"
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