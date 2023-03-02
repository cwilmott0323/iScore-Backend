package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type City struct {
	CityId        int64  `gorm:"primary_key;auto_increment" json:"city_id"`
	CityName      string `gorm:"size(255);not null;" json:"city_name"`
	ImageLocation string `gorm:"size(255);not null;" json:"image_location"`
	CountryId     int64  `gorm:"" json:"country_id"`
}

type CityReturn struct {
	CityId        int64  `gorm:"primary_key;auto_increment" json:"city_id"`
	CityName      string `gorm:"size(255);not null;" json:"city_name"`
	ImageLocation string `gorm:"size(255);not null;" json:"image_location"`
	CountryId     int64  `gorm:"" json:"country_id"`
	CountryName   string `gorm:"size(255);not null;" json:"country_name"`
}

func (c *City) GetCities(db *gorm.DB, countryName string) ([]CityReturn, error) {
	var err error
	var x []CityReturn

	// select ci.city_id, ci.city_name, ci.image_location, co.country_id, co.country_name from cities AS ci inner join country_city_activity_maps AS map on ci.city_id = map.city_id inner join countries AS co on map.country_id = co.country_id;
	err = db.Debug().Table("cities as ci").Select("DISTINCT ci.city_id, ci.city_name, ci.image_location, co.country_id, co.country_name").Joins("left join country_city_activity_maps AS map on ci.city_id = map.city_id").Joins("left join countries AS co on map.country_id = co.country_id").Where("co.country_name = ?", countryName).Find(&x).Error

	if len(x) == 0 {
		return []CityReturn{}, errors.New("no Cities Returned")
	}

	if err != nil {
		return []CityReturn{}, err
	}

	return x, err
}
