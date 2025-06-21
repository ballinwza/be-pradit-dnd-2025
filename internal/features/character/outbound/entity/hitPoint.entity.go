package character_entity

type HitPointEntity struct {
	MaxHp       int32 `bson:"max_hp"`
	CurrentHp   int32 `bson:"current_hp"`
	TemporaryHp int32 `bson:"temporary_hp"`
}
