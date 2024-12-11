package updatebal

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(cust *customer, tranc transaction) error {
	if tranc.transactionType != transactionDeposit && tranc.transactionType != transactionWithdrawal { 
		return errors.New("unknown transaction type")
	}
	
	if tranc.transactionType == transactionDeposit {
		*cust = customer{
			id: tranc.customerID,
			balance: cust.balance + tranc.amount,
		}
	} 

	if tranc.transactionType == transactionWithdrawal {
		if cust.balance < tranc.amount {
			return errors.New("insufficient funds")	
		}
		*cust = customer{
			id: tranc.customerID,
			balance: cust.balance - tranc.amount,
		}
	}

	return nil 
}
