package ability_detail_service

import ability_detail_outbound_repository "github.com/ballinwza/be-pradit-dnd-2025/internal/features/ability_detail/outbound/repository"

type AbilityDetailService struct {
	abilityDetailRepository *ability_detail_outbound_repository.AbilityDetailRepository
}

func NewAbilityDetailService() *AbilityDetailService {
	abilityDetailRepository := ability_detail_outbound_repository.NewAbilityDetailRepository()

	return &AbilityDetailService{
		abilityDetailRepository: abilityDetailRepository,
	}
}