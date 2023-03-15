package models

import (
	"errors"
	"gorm.io/gorm"

	"strconv"
)

type Activity struct {
	ActivityID    int64   `gorm:"primary_key;auto_increment" json:"activity_id"`
	ActivityName  string  `gorm:"size(255);not null;unique" json:"activity_name"`
	ImageLocation string  `gorm:"size(255);not null;" json:"image_location"`
	ActivityType  string  `gorm:"size(255);not null;" json:"activity_type"`
	Sponsored     bool    `gorm:"" json:"sponsored"`
	Points        int64   `gorm:"" json:"points"`
	LatX          float64 `gorm:"'" json:"latx"`
	LatY          float64 `gorm:"'" json:"latY"`
	LonX          float64 `gorm:"" json:"lonx"`
	LonY          float64 `gorm:"" json:"lony"`
}

type UserActivityMap struct {
	ActivityID int64
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

	// select * from activities inner join country_city_activity_maps on activities.activity_id = country_city_activity_maps.activity_id inner join cities on country_city_activity_maps.city_id = cities.city_id inner join countries on country_city_activity_maps.country_id = countries.country_id
	err = db.Debug().Table("activities as a").Select("a.activity_id, a.activity_name, a.image_location, a.activity_type, a.sponsored, a.points, a.lat_x, a.lat_y, a.lon_x, a.lon_y").Joins("left join country_city_activity_maps as map on a.activity_id = map.activity_id").Joins("left join cities as ci on map.city_id = ci.city_id").Joins("left join countries as co on map.country_id = co.country_id").Where("ci.city_name = ?", cityName).Find(&x).Error

	if len(x) == 0 {
		return []Activity{}, errors.New("no Activity Returned")
	}

	return x, err
}

func (c *Activity) GetActivityLocation(db *gorm.DB, activity string) ([]Activity, error) {
	var err error
	var x []Activity

	// select * from activities inner join country_city_activity_maps on activities.activity_id = country_city_activity_maps.activity_id inner join cities on country_city_activity_maps.city_id = cities.city_id inner join countries on country_city_activity_maps.country_id = countries.country_id
	err = db.Debug().Table("activities").Select("points, lat_x,lat_y, lon_x, lon_y").Where("activity_id = ?", activity).First(&x).Error

	if len(x) == 0 {
		return []Activity{}, errors.New("no Activity Returned")
	}

	return x, err
}

func (c *Account) AddPointsDB(db *gorm.DB, userID uint32, points int64) (Account, error) {
	var err error
	var x Account
	// select * from activities inner join country_city_activity_maps on activities.activity_id = country_city_activity_maps.activity_id inner join cities on country_city_activity_maps.city_id = cities.city_id inner join countries on country_city_activity_maps.country_id = countries.country_id

	db.Debug().Model(&Account{}).Where("account_id = ?", userID).Update("points", gorm.Expr("points + ?", points))

	return x, err
}

func (c *Account) CompleteActivity(db *gorm.DB, userID uint32, activity int64) (Account, error) {
	var err error
	var x Account
	// select * from activities inner join country_city_activity_maps on activities.activity_id = country_city_activity_maps.activity_id inner join cities on country_city_activity_maps.city_id = cities.city_id inner join countries on country_city_activity_maps.country_id = countries.country_id

	y := AccountActivity{ActivityId: activity, Completed: true}

	db.Debug().Table("account_" + strconv.FormatInt(int64(userID), 10) + "_activities").Create(&y)

	return x, err
}

func (c *Activity) GetAllActivities(db *gorm.DB) ([]Activity, error) {
	var err error
	var x []Activity

	// select * from activities inner join country_city_activity_maps on activities.activity_id = country_city_activity_maps.activity_id inner join cities on country_city_activity_maps.city_id = cities.city_id inner join countries on country_city_activity_maps.country_id = countries.country_id
	err = db.Debug().Table("activities").Select("activity_id").Find(&x).Error

	if len(x) == 0 {
		return []Activity{}, errors.New("no Activity Returned")
	}

	return x, err
}

func (c *Activity) FillAccountActivities(m []Activity, db *gorm.DB, account_id int64) ([]Activity, error) {
	var err error
	var x []Activity
	var y []AccountActivity

	for _, v := range m {

		y = append(y, []AccountActivity{{ActivityId: v.ActivityID}}...)
	}

	db.Debug().Table("account_"+strconv.FormatInt(account_id, 10)+"_activities").CreateInBatches(&y, 100)
	return x, err
}

func (c *Activity) IsComplete(db *gorm.DB, userID uint32, activity string) ([]AccountActivity, error) {
	var err error
	var x []AccountActivity

	// select completed from account_38_activities inner join activities on account_38_activities.activity_id = activities.activity_id where activities.activity_id = 1;
	err = db.Debug().Table("account_"+strconv.Itoa(int(userID))+"_activities").Select("*").Joins("inner join activities on account_"+strconv.Itoa(int(userID))+"_activities.activity_id = activities.activity_id").Where("activities.activity_id = ?", activity).First(&x).Error

	if len(x) == 0 {

		return nil, nil
	}

	return x, err
}
