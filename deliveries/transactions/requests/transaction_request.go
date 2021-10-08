package requests

import (
	"mikro/business/transactions"
)

type TransactionRequest struct {
	ProductId int `json:"product_id"`
	BuyerId   int `json:"buyer_id"`
}

func (ur *TransactionRequest) ToDomainUser() transactions.Domain {
	return transactions.Domain{
		ProductId: ur.ProductId,
		BuyerId:   ur.BuyerId,
	}
}
