package core_outbound_entity

type CoinShortType string

const (
	CoinShortTypeCp CoinShortType = "cp"
	CoinShortTypeSp CoinShortType = "sp"
	CoinShortTypeEp CoinShortType = "ep"
	CoinShortTypeGp CoinShortType = "gp"
	CoinShortTypePp CoinShortType = "pp"
)

type CoinEntity struct {
	Name      string        `bson:"name" json:"name"`
	Value     float64       `bson:"value" json:"value"`
	ShortType CoinShortType `bson:"short_type" json:"short_type"`
}
