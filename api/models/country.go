package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Country struct {
	CountryId     int64  `gorm:"primary_key;auto_increment" json:"country_id"`
	CountryName   string `gorm:"size(255);not null;unique;" json:"country_name"`
	ImageLocation string `gorm:"size(255);not null;" json:"image_location"`
}

func (c *Country) GetCountries(db *gorm.DB) ([]Country, error) {
	var err error
	var x []Country
	err = db.Debug().Find(&x).Error

	fmt.Println("Get Countries: ", x)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []Country{}, errors.New("no Cities Returned")
	}

	if err != nil {
		return []Country{}, err
	}

	return x, err
}
