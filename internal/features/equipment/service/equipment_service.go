package equipment_service

import (
	equipment_outbound_repository "github.com/ballinwza/be-pradit-dnd-2025/internal/features/equipment/outbound/repository"
	weapon_outbound_repository "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/outbound/repository"
)

type EquipmentService struct {
	equipmentRepository *equipment_outbound_repository.EquipmentRepository
	weaponRepository    *weapon_outbound_repository.WeaponRepository
}

func NewEquipmentService() *EquipmentService {
	equipmentRepository := equipment_outbound_repository.NewEquipmentRepository()
	weaponRepository := weapon_outbound_repository.NewWeaponRepository()
	return &EquipmentService{
		equipmentRepository: equipmentRepository,
		weaponRepository:    weaponRepository,
	}
}
