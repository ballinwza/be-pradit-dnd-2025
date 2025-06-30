package character_outbound_entity

type CharacterPocketMoneyEntity struct {
	Cp float64 `bson:"cp"`
	Sp float64 `bson:"sp"`
	Ep float64 `bson:"ep"`
	Gp float64 `bson:"gp"`
	Pp float64 `bson:"pp"`
}
