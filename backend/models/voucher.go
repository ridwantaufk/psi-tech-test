package models

type Voucher struct {
	ID       string  `gorm:"primaryKey;type:varchar(50)" json:"id"`
	Code     string  `gorm:"unique" json:"code"`
	Discount float64 `json:"discount"`
	IsActive bool    `json:"is_active"`
}
