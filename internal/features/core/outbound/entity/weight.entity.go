package core_outbound_entity

type WeightEntity struct {
	Unit  string `bson:"unit" json:"unit"`
	Value int32  `bson:"value" json:"value"`
}
