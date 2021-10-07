package records

type Transactions struct {
	Id        int `gorm:"primaryKey"`
	ProductId int
	BuyerId   int
	Product   Products `gorm:"foreignKey:product_id"`
	Buyer     Users    `gorm:"foreignKey:buyer_id"`
	Status    bool
}
