package character_outbound_entity

type AbilityEntity struct {
	Str *int32 `bson:"str"`
	Dex *int32 `bson:"dex"`
	Con *int32 `bson:"con"`
	Int *int32 `bson:"int"`
	Wis *int32 `bson:"wis"`
	Cha *int32 `bson:"cha"`
}
