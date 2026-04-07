package models

type Company struct {
	ID          string `gorm:"primaryKey;type:varchar(50)" json:"id"`
	UserID      string `json:"user_id"`
	CompanyCode string `json:"company_code"`
	CompanyName string `json:"company_name"`
	User        User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
