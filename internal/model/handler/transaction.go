package handler

import "net/http"

type TransactionHandler interface {
	WithdrawHandler(w http.ResponseWriter, r *http.Request)
}

type WithdrawReqBody struct {
	AccountID  string  `json:"account_id"`
	Amount     float64 `json:"amount"`
	Withdrawer string  `json:"withdrawer"`
}
