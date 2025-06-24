package character_outbound_entity

type AbilityEntity struct {
	Strength     int32 `bson:"strength" `
	Dexterity    int32 `bson:"dexterity" `
	Constitution int32 `bson:"constitution" `
	Intelligence int32 `bson:"intelligence" `
	Wisdom       int32 `bson:"wisdom" `
	Charisma     int32 `bson:"charisma" `
}
