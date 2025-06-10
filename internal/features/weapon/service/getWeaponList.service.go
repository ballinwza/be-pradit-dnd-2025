package weapon_service

import (
	"context"
	"log"

	weapon_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func (s *WeaponService) GetWeaponList(ctx context.Context) ([]*model.Weapon, error) {
	weaponList, err := s.weaponRepository.FindAll(ctx)
	if err != nil {
		log.Println("WeaponService.GetWeaponList Error : ", err)
		return nil, err
	}

	result := make([]*model.Weapon, len(weaponList))
	for i, weapon := range weaponList {
		afterMapping := weapon_mapper.MapperWeaponEntityToModel(*weapon)
		result[i] = &afterMapping
	}

	return result, nil
}
