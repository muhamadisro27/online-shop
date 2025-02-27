package entity

type Product struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;type:serial;column:id"`
	Name     string `gorm:"type:varchar(255);column:name"`
	Price    int    `gorm:"type:int;column:price"`
	IsDelete bool `gorm:"type:b"`
}

func (Product) TableName() string {
	return "products"
}
