package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Account struct {
	AccountId int64     `gorm:"primary_key;auto_increment" json:"id"`
	Password  string    `gorm:"size(255);not null;" json:"password"`
	Email     string    `gorm:"size(255);not null;unique" json:"email"`
	Name      string    `gorm:"size(255);not null;" json:"name"`
	Points    int       `gorm:"not null" json:"points"`
	CreatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_on"`
	LastLogin time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"last_login"`
}

func (a *Account) FindAccountByID(db *gorm.DB, aid uint32) (*Account, error) {
	var err error
	err = db.Debug().Model(Account{}).Select("name, points").Where("account_id = ?", aid).Take(&a).Error
	if gorm.IsRecordNotFoundError(err) {
		return &Account{}, errors.New("Account Not Found")
	}
	if err != nil {
		return &Account{}, err
	}
	return a, err
}

func (a *Account) CreateAccount(db *gorm.DB) (*Account, error) {
	var err error
	err = db.Debug().Create(&a).Error
	if err != nil {
		return &Account{}, err
	}
	return a, err
}
