package response

type WithdrawReply struct {
	AccountID      string
	UserName       string
	WithdrawAmount float64
	RemainAmount   float64
}
