package models

import "time"

type AuthUser struct {
	ID        string    `gorm:"primaryKey;type:varchar(50)" json:"id"`
	Username  string    `gorm:"unique" json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
