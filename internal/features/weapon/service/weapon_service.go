package weapon_service

import weapon_outbound_repository "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/outbound/repository"

type WeaponService struct {
	weaponRepository *weapon_outbound_repository.WeaponRepository
}

func NewWeaponService() *WeaponService {
	weaponRepository := weapon_outbound_repository.NewWeaponRepository()

	return &WeaponService{
		weaponRepository: weaponRepository,
	}
}
