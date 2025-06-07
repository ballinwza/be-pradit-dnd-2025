package armor_service

import armor_outbound_repository "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/outbound/repository"

type ArmorService struct {
	armorRepository *armor_outbound_repository.ArmorRepository
}

func NewArmorService() *ArmorService {
	armorRepository := armor_outbound_repository.NewArmorRepository()

	return &ArmorService{
		armorRepository: armorRepository,
	}
}