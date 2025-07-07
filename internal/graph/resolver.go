package graph

import (
	"context"
	"errors"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/error_handler"
	ability_detail_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/ability_detail/service"
	armor_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/service"
	cache_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/cache/service"
	character_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/service"
	class_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/class/service"
	equipment_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/equipment/service"
	level_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/level/service"
	user_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/service"
	weapon_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/service"
	"github.com/ballinwza/be-pradit-dnd-2025/middleware"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService          *user_service.UserService
	CharacterService     *character_service.CharacterService
	AbilityDetailService *ability_detail_service.AbilityDetailService
	ArmorService         *armor_service.ArmorService
	WeaponService        *weapon_service.WeaponService
	EquipmentService     *equipment_service.EquipmentService
	ClassService         *class_service.ClassService
	LevelService         *level_service.LevelService
	CacheService         *cache_service.CacheService
}

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	claims, ok := ctx.Value(middleware.UserCtxKey).(*middleware.UserClaims)

	if !ok || claims == nil {
		return nil, error_handler.NewValidationError("access denied: you must be logged in", errors.New("invalid token"), http.StatusInternalServerError).GqlError(ctx)
	}

	return next(ctx)
}
