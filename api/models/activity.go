package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Activity struct {
	ActivityID    int64  `gorm:"primary_key;auto_increment" json:"activity_id"`
	ActivityName  string `gorm:"size(255);not null;" json:"activity_name"`
	ImageLocation string `gorm:"size(255);not null;" json:"image_location"`
	ActivityType  string `gorm:"size(255);not null;" json:"activity_type"`
	Sponsored     bool   `gorm:"" json:"sponsored"`
	Points        int64  `gorm:"" json:"points"`
}

func (c *Activity) GetCitiesInfo(db *gorm.DB, countryName string, cityName string) ([]Activity, error) {
	var err error
	var x []Activity

	err = db.Debug().Table("city_infos").Select("*").Where("city_name ="+
		" ?", cityName).Where("country_name = ?", countryName).Find(&x).Error

	if len(x) == 0 {
		return []Activity{}, errors.New("no Cities Returned")
	}

	return x, err
}

func (c *Activity) GetActivityInfo(db *gorm.DB, countryName string, cityName string, activityName string) ([]Activity, error) {
	var err error
	var x []Activity

	err = db.Debug().Table("activities").Select("*").Where("activity_name = ?", activityName).First(&x).Error

	if len(x) == 0 {
		return []Activity{}, errors.New("no Activity Returned")
	}

	return x, err
}

func (c *Activity) GetActivities(db *gorm.DB, countryName string, cityName string) ([]Activity, error) {
	var err error
	var x []Activity

	fmt.Println("GET ACTIVITES!!")
	// select * from activities inner join country_city_activity_maps on activities.activity_id = country_city_activity_maps.activity_id inner join cities on country_city_activity_maps.city_id = cities.city_id inner join countries on country_city_activity_maps.country_id = countries.country_id
	err = db.Debug().Table("activities as a").Select("*").Joins("left join country_city_activity_maps as map on a.activity_id = map.activity_id").Joins("left join cities as ci on map.city_id = ci.city_id").Joins("left join countries as co on map.country_id = co.country_id").Where("ci.city_name = ?", cityName).Find(&x).Error

	if len(x) == 0 {
		return []Activity{}, errors.New("no Activity Returned")
	}

	return x, err
}
