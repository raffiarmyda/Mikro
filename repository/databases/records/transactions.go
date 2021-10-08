package records

import "mikro/business/transactions"

type Transactions struct {
	Id        int `gorm:"primaryKey"`
	ProductId int
	BuyerId   int
	Product   Products `gorm:"foreignKey:product_id"`
	Buyer     Users    `gorm:"foreignKey:buyer_id"`
}

func (p *Transactions) TransactionsToDomain() transactions.Domain {
	return transactions.Domain{
		ID:        p.Id,
		ProductId: p.ProductId,
		BuyerId:   p.BuyerId,
		Product: transactions.Products(struct {
			Id    int
			Name  string
			Price float64
		}{Id: p.Product.ID, Name: p.Product.Name, Price: p.Product.Price}),
		Buyer: struct {
			ID          int
			Name        string
			StoreName   string
			City        string
			Phone       string
			Username    string
			BankAccount string
			NoAccount   string
			Password    string
		}{ID: p.Buyer.ID, Name: p.Buyer.Name, StoreName: p.Buyer.StoreName, City: p.Buyer.City, Phone: p.Buyer.Phone, Username: p.Buyer.Username, BankAccount: p.Buyer.BankAccount, NoAccount: p.Buyer.NoAccount, Password: p.Buyer.Password},
	}
}

func TransactionsToListDomain(data []Transactions) []transactions.Domain {
	var list []transactions.Domain
	for _, v := range data {
		list = append(list, v.TransactionsToDomain())
	}
	return list
}

func TransactionsFromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		Id:        domain.ID,
		ProductId: domain.ProductId,
		BuyerId:   domain.BuyerId,
		//Product:   domain.Product,
		//Buyer:     domain.Buyer,
	}
}
