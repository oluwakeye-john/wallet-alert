package models

import (
	"errors"
	"log"

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

func (a *Account) GetWithAddress(db *gorm.DB) (*Address, error) {
	addresses := []Address{}
	result := db.Joins("address").First(a, "email = ?", a.Email)

	if result.Error != nil {
		return &Address{}, result.Error
	}

	if len(addresses) < 1 {
		return &Address{}, errors.New("Invalid")
	}
	return &addresses[0], nil
}

func (a *Account) IncrementTransactionCount(db *gorm.DB) *gorm.DB {
	db.Transaction(func(tx *gorm.DB) error {
		if err := a.Get(tx).Error; err != nil {
			return err
		}
		if err := tx.Update("transaction_count", a.TransactionCount+1).Error; err != nil {
			return err
		}
		return nil
	})
	return db
}

func (a *Account) AfterDelete(tx *gorm.DB) error {
	log.Println("ZHere")
	address := Address{}
	address.ID = a.AddressId

	err := address.DeleteAddressWithNoAccount(tx)
	if err != nil {
		log.Println(err)
	}

	return nil
}
