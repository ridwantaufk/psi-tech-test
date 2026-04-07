package models

import "github.com/google/uuid"

type User struct {
	ID    string `gorm:"primaryKey;type:varchar(50)" json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Telp  string `json:"telp"`
}

func (u *User) BeforeCreate(tx interface{}) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}
