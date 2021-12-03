package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Address      string `json:"address" gorm:"not null"`
	Email        string `json:"email" gorm:"not null"`
	CurrencyCode string `json:"currency_code" gorm:"not null"`
}

func (a *Account) Create(db *gorm.DB) *gorm.DB {
	return db.FirstOrCreate(a, "email = ? AND address = ? AND currency_code = ?", a.Email, a.Address, a.CurrencyCode)
}

func (a *Account) Delete(db *gorm.DB) *gorm.DB {
	return db.Delete(a, "address = ? AND email = ?", a.Address, a.Email)
}

func (a *Account) Get(db *gorm.DB) *gorm.DB {
	return db.First(a, "email = ? AND address = ?", a.Email, a.Address)
}
