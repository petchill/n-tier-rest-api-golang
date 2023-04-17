package bank

type Account struct {
	ID           string  `json:"id" bson:"id"`
	UserID       int     `json:"user_id" bson:"user_id"`
	AmountRemain float64 `json:"amount_remain" bson:"amount_remain"`
	IsAvailable  bool    `json:"is_available" bson:"is_available"`
}
