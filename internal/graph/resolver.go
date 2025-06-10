package graph

import (
	ability_detail_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/ability_detail/service"
	armor_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/service"
	character_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/service"
	user_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/service"
	weapon_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/service"
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
}
