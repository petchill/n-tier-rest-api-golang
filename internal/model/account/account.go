package bank

type Account struct {
	ID               string  `json:"id"`
	UserID           int     `json:"user_id"`
	AmountRemain     float64 `json:"remain_amount"`
	AccountAvailable bool    `json:"account_avaliable"`
}
