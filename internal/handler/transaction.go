package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	mHandler "github.com/petchill/n-tier-rest-api-golang/internal/model/handler"
	mService "github.com/petchill/n-tier-rest-api-golang/internal/model/service"
)

type transactionHandler struct {
	transactionService mService.TransactionService
}

// WithdrawHandler implements handler.TransactionHandler
func (h transactionHandler) WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	// allow only POST Method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	withdrawPayload := mHandler.WithdrawReqBody{}
	// bind req body to WithdrawReqBody struct
	err := json.NewDecoder(r.Body).Decode(&withdrawPayload)
	if err != nil {
		http.Error(w, "Some request information are missing.", http.StatusBadRequest)
		return
	}

	if bodyValidateErr := validateRequestBody(withdrawPayload); bodyValidateErr != nil {
		http.Error(w, bodyValidateErr.Error(), http.StatusBadRequest)
		return
	}

	// call service function
	reply, err := h.transactionService.Withdraw(r.Context(), withdrawPayload)
	if err != nil {
		errMsg := fmt.Sprintf("Withdraw process is fail because of %s", err.Error())
		http.Error(w, errMsg, http.StatusUnprocessableEntity)
		return
	}

	replyMsg := fmt.Sprintf("%s withdraw money %v baht from account %s success. Remained %v baht.", reply.UserName, reply.WithdrawAmount, reply.AccountID, reply.RemainAmount)
	// reply message to client
	fmt.Fprintf(w, replyMsg)
	return
}

func validateRequestBody(withdrawPayload mHandler.WithdrawReqBody) error {
	if nameValid := validateEnglishName(withdrawPayload.Withdrawer); !nameValid {
		return errors.New("withdrawer must be English name and in format (fisrtname surname).")
	}

	if accountIDValid := validateAccountID(withdrawPayload.AccountID); !accountIDValid {
		return errors.New("account_id must be number with 8 length.")
	}

	if amountValid := validateAmount(withdrawPayload.Amount); !amountValid {
		return errors.New("withdraw_amount must be more than 0.")
	}

	return nil
}

func validateEnglishName(name string) bool {
	regex, _ := regexp.Compile("^[a-zA-Z]+ [a-zA-Z]+$")
	return regex.MatchString(name)
}

func validateAccountID(id string) bool {
	// Assume as id must have 8 length numbers
	regex, _ := regexp.Compile("^[0-9]{8}$")
	return regex.MatchString(id)
}

func validateAmount(amount float64) bool {
	return amount > 0
}

func NewTransactionHandler(transactionService mService.TransactionService) mHandler.TransactionHandler {
	return transactionHandler{
		transactionService: transactionService,
	}
}
