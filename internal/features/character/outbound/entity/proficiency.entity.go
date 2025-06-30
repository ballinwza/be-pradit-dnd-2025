package character_outbound_entity

type CharacterProficiencyEntity struct {
	Athletics      *int32 `bson:"athletics"`
	Arobatics      *int32 `bson:"arobatics"`
	SleightOfHand  *int32 `bson:"sleight_of_hand"`
	Stealth        *int32 `bson:"stealth"`
	Arcana         *int32 `bson:"arcana"`
	History        *int32 `bson:"history"`
	Investigation  *int32 `bson:"investigation"`
	Nature         *int32 `bson:"nature"`
	Religion       *int32 `bson:"religion"`
	AnimalHandling *int32 `bson:"animal_handling"`
	Insight        *int32 `bson:"insight"`
	Medicine       *int32 `bson:"medicine"`
	Perception     *int32 `bson:"perception"`
	Survival       *int32 `bson:"survival"`
	Deception      *int32 `bson:"deception"`
	Intimidation   *int32 `bson:"intimidation"`
	Performance    *int32 `bson:"performance"`
	Persuasion     *int32 `bson:"persuasion"`
}
