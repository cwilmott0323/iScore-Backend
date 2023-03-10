package models

import (
	"gorm.io/gorm"
	"log"
)

var Countries = []Country{
	{
		CountryId:     1,
		CountryName:   "England",
		ImageLocation: "test",
	},
}

var Cities = []City{
	{
		CityId:        1,
		CityName:      "London",
		ImageLocation: "test",
	},
}

var CityInfos = []Activity{
	{
		ActivityID:    1,
		ActivityName:  "Big Ben",
		ImageLocation: "test",
		ActivityType:  "Place",
		Sponsored:     false,
		Points:        10,
	},
}

func TestData(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&Account{}, &City{}, &Activity{}, &Country{}, &CountryCityActivityMap{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&Account{}, &City{}, &Activity{}, &Country{}, &CountryCityActivityMap{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range Countries {
		results := db.Model(&Country{}).Create(&Countries[i])

		if results.Error != nil {
			log.Fatalf("cannot seed cards table: %v", err)
		}

	}
	for i, _ := range Cities {
		results := db.Debug().Model(&City{}).Create(&Cities[i]).Error
		if results.Error != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range CityInfos {
		results := db.Debug().Model(&Activity{}).Create(&CityInfos[i]).Error
		if results.Error != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}

func Load(db *gorm.DB) {

	err := db.AutoMigrate(&Account{}, &City{}, &Activity{}, &Country{}, &CountryCityActivityMap{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}
