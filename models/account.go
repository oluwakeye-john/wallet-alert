package models

import (
	"log"

	"github.com/oluwakeye-john/wallet-alert/blockcypher"
	"gorm.io/gorm"
)

type Account struct {
	Base
	Email            string `json:"email" gorm:"unique;not null"`
	TransactionCount int    `json:"transaction_count"`
	Address          string `json:"address" gorm:"unique;not null"`
	HookId           string `json:"hook_id"`
	CurrencyCode     string `json:"currency_code"`
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

func (a *Account) GetByAddress(db *gorm.DB) *gorm.DB {
	return db.First(a, "address = ?", a.Address)
}

func (a *Account) IncrementTransactionCount(db *gorm.DB) error {
	if err := db.Model(a).Update("transaction_count", a.TransactionCount+1).Error; err != nil {
		return err
	}
	return nil
}

func (a *Account) SetupHook(tx *gorm.DB) (err error) {
	hook, error := blockcypher.SetupAddressTransactionHook(a.Address, a.CurrencyCode)
	if error != nil {
		log.Println(error)
		return error
	}

	result := tx.Model(a).UpdateColumn("hook_id", hook.ID)

	if result.Error != nil {
		log.Println(err)
		return
	}

	return
}

func (a *Account) DeleteHook(db *gorm.DB) error {
	if error := blockcypher.DeleteAddressTransactionHook(a.HookId, a.CurrencyCode); error != nil {
		return error
	}

	return nil
}
