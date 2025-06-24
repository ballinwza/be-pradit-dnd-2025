package character_outbound_entity

type ProficiencyEntity struct {
	Athletics      int32 `bson:"athletics" json:"athletics"`
	Arobatics      int32 `bson:"arobatics" json:"arobatics"`
	SleightOfHand  int32 `bson:"sleight_of_hand" json:"sleight_of_hand"`
	Stealth        int32 `bson:"stealth" json:"stealth"`
	Arcana         int32 `bson:"arcana" json:"arcana"`
	History        int32 `bson:"history" json:"history"`
	Investigation  int32 `bson:"investigation" json:"investigation"`
	Nature         int32 `bson:"nature" json:"nature"`
	Religion       int32 `bson:"religion" json:"religion"`
	AnimalHandling int32 `bson:"animal_handling" json:"animal_handling"`
	Insight        int32 `bson:"insight" json:"insight"`
	Medicine       int32 `bson:"medicine" json:"medicine"`
	Perception     int32 `bson:"perception" json:"perception"`
	Survival       int32 `bson:"survival" json:"survival"`
	Deception      int32 `bson:"deception" json:"deception"`
	Intimidation   int32 `bson:"intimidation" json:"intimidation"`
	Performance    int32 `bson:"performance" json:"performance"`
	Persuasion     int32 `bson:"persuasion" json:"persuasion"`
}
