package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email            string `json:"email" gorm:"not null"`
	TransactionCount int    `json:"transaction_count"`
	AddressId        uint   `json:"address_id" gorm:"not null"`
}

func (a *Account) Create(db *gorm.DB) *gorm.DB {
	return db.Create(a)
}

func (a *Account) Delete(db *gorm.DB) *gorm.DB {
	return db.Delete(a, "email = ?", a.Email)
}

func (a *Account) Get(db *gorm.DB) *gorm.DB {
	return db.First(a, "email = ?", a.Email)
}

func (a *Account) IncrementTransactionCount(db *gorm.DB) error {
	if err := db.Model(a).Update("transaction_count", a.TransactionCount+1).Error; err != nil {
		return err
	}
	return nil
}
