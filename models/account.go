package models

import (
	"github.com/oluwakeye-john/wallet-alert/blockcypher"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Address      string `json:"address" gorm:"not null"`
	Email        string `json:"email" gorm:"not null"`
	CurrencyCode string `json:"currency_code" gorm:"not null"`
	HookId       string `json:"hook_id"`
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

func (a *Account) BeforeDelete(tx *gorm.DB) (err error) {
	blockcypher.DeleteHookOnAddress(a.HookId, a.CurrencyCode)
	return
}
