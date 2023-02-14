package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type City struct {
	CityId        int64  `gorm:"primary_key;auto_increment" json:"city_id"`
	CityName      string `gorm:"size(255);not null;" json:"city_name"`
	ImageLocation string `gorm:"size(255);not null;" json:"image_location"`
	CountryId     int64  `gorm:"" json:"country_id"`
}

type CityInfo struct {
	ActivityID    int64  `gorm:"primary_key;auto_increment" json:"activity_id"`
	ActivityName  string `gorm:"size(255);not null;" json:"activity_name"`
	CityId        int64  `gorm:"" json:"city_id"`
	CityName      string `gorm:"size(255);not null;" json:"city_name"`
	ImageLocation string `gorm:"size(255);not null;" json:"image_location"`
	CountryId     int64  `gorm:"" json:"country_id"`
	CountryName   string `gorm:"size(255);not null;" json:"country_name"`
	ActivityType  string `gorm:"size(255);not null;" json:"activity_type"`
	Sponsored     bool   `gorm:"" json:"sponsored"`
	Points        int64  `gorm:"" json:"points"`
}

func (c *City) GetCities(db *gorm.DB, countryName string) ([]City, error) {
	var err error
	var x []City

	err = db.Debug().Table("cities").Select("*").Joins("left join countries on cities.country_id = countries.country_id").Where("countries.country_name = ?", countryName).Find(&x).Error

	if gorm.IsRecordNotFoundError(err) {
		return []City{}, errors.New("no Cities Returned")
	}
	if err != nil {
		return []City{}, err
	}
	fmt.Println(x)
	return x, err
}

func (c *CityInfo) GetCitiesInfo(db *gorm.DB, countryName string, cityName string) ([]CityInfo, error) {
	var err error
	var x []CityInfo

	err = db.Debug().Table("city_infos").Select("*").Where("city_name = ?", cityName).Where("country_name = ?", countryName).Find(&x).Error

	if len(x) == 0 {
		log.Println("No records found")
		return []CityInfo{}, errors.New("no Cities Returned")
	}
	//fmt.Println("gorm Error: ", err)
	//
	//if gorm.IsRecordNotFoundError(err) {
	//	log.Println("No records found")
	//	return []CityInfo{}, errors.New("no Cities Returned")
	//}
	//if err != nil {
	//	return []CityInfo{}, err
	//}
	fmt.Println(x)
	return x, err
}
