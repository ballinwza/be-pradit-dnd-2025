package ability_detail_service

import (
	"context"
	"log"

	ability_detail_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/ability_detail/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func (s *AbilityDetailService) GetAbilityDetailByShort(ctx context.Context, short model.AbilityShortType) (*model.AbilityDetail, error){

	abilityDetail, err := s.abilityDetailRepository.FindOneByShort(context.Background(), ability_detail_mapper.MapperAbilityDetailShortTypeModelToEntity(short))
	if err != nil {
		log.Printf("Error fetching abilityDetail: %v", err)
	}

	result := ability_detail_mapper.MapperAbilityDetailEntityToModel(*abilityDetail)
	return &result, nil
}