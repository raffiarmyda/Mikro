package responses

import (
	"mikro/business/transactions"
)

type TransactionResponse struct {
	Id        int         `json:"id"`
	ProductId int         `json:"product_id"`
	BuyerId   int         `json:"buyer_id"`
	Product   interface{} `json:"product"`
	Buyer     interface{} `json:"buyer"`
}

func FromProductDomain(domain transactions.Domain) TransactionResponse {
	return TransactionResponse{
		Id:        domain.ID,
		ProductId: domain.ProductId,
		BuyerId:   domain.BuyerId,
		Product:   domain.Product,
		Buyer:     domain.Buyer,
	}
}

func FromProductListDomain(listDomain []transactions.Domain) []TransactionResponse {
	var list []TransactionResponse
	for _, v := range listDomain {
		list = append(list, FromProductDomain(v))
	}
	return list
}
