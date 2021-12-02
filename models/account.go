package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Address      string `json:"address" gorm:"unique"`
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	CurrencyCode string `json:"currency_code"`
}

func (a *Account) Create(db *gorm.DB) *gorm.DB {
	return db.Create(a)
}

func (a *Account) Delete(db *gorm.DB) *gorm.DB {
	return db.Delete(a)
}

func (a *Account) Get(db *gorm.DB) *gorm.DB {
	return db.First(a)
}
