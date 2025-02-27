package entity

type Product struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;type:serial;column:id"`
	Name     string `gorm:"type:varchar(255);column:name"`
	Category string `gorm:"type:varchar(50);column:category"`
	Price    int    `gorm:"type:int;column:price"`
}

func (Product) TableName() string {
	return "product"
}
