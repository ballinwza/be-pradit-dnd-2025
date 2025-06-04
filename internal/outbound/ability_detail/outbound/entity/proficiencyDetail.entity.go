package ability_detail_outbound_entity



type ProficiencyDetailEntity struct {
	Name string `bson:"name" json:"name"`
	DescriptionEn string `bson:"description_en" json:"description_en"`
	DescriptionTh string `bson:"description_th" json:"description_th"`
}
