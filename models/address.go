package models

import (
	"log"

	"github.com/oluwakeye-john/wallet-alert/blockcypher"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Address      string    `json:"address" gorm:"not null"`
	CurrencyCode string    `json:"currency_code" gorm:"not null"`
	HookId       string    `json:"hook_id" gorm:"not null"`
	Accounts     []Account `json:"accounts"`
}

func (addr *Address) GetOrCreate(db *gorm.DB) *gorm.DB {
	return db.FirstOrCreate(addr, "address = ? AND currency_code = ?", addr.Address, addr.CurrencyCode)
}

func (addr *Address) Get(db *gorm.DB) *gorm.DB {
	return db.First(addr)
}

func (addr *Address) DeleteAddressWithNoAccount(db *gorm.DB) error {
	result := db.Preload("Accounts").FirstOrCreate(addr, "id = ?", addr.ID)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Number of accounts", len(addr.Accounts))
	if len(addr.Accounts) == 0 {
		if error := blockcypher.DeleteAddressTransactionHook(addr.HookId, addr.CurrencyCode); error != nil {
			return error
		}

		if error := db.Delete(addr, "id = ?", addr.ID).Error; error != nil {
			return error
		}

		return nil
	}
	log.Println("Contains other accounts")
	return nil
}

func (addr *Address) BeforeCreate(tx *gorm.DB) (err error) {
	hook, error := blockcypher.SetupAddressTransactionHook(addr.Address, addr.CurrencyCode)
	if error != nil {
		return error
	}

	addr.HookId = hook.ID
	return
}
