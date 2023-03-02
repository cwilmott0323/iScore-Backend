package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

//var card_list = []models.Card{
//	{
//		CardName:    "BlueEyes",
//		Attack:      3000,
//		Defence:     2500,
//		Description: "White Dragon",
//		Ability:     "",
//		Typing:      "Dragon",
//		Rarity:      "Rare",
//		Type:        "Dragon",
//		Set:         "Legend of Blue Eyes White Dragon",
//		SetCode: "LOB",
//		},
//	//models.Card{
//	//	CardName: "Dark Magician",
//	//},
//}

//type Country struct {
//	CountryId     int64  `gorm:"primary_key;auto_increment" json:"country_id"`
//	CountryName   string `gorm:"size(255);not null;" json:"country_name"`
//	ImageLocation string `gorm:"size(255);not null;" json:"image_location"`
//}

var Countries = []Country{
	{
		CountryId:     1,
		CountryName:   "England",
		ImageLocation: "test",
	},
}

//type City struct {
//	CityId        int64  `gorm:"primary_key;auto_increment" json:"city_id"`
//	CityName      string `gorm:"size(255);not null;" json:"city_name"`
//	ImageLocation string `gorm:"size(255);not null;" json:"image_location"`
//	CountryId     int64  `gorm:"" json:"country_id"`
//}

var Cities = []City{
	{
		CityId:        1,
		CityName:      "London",
		ImageLocation: "test",
		CountryId:     1,
	},
}

//type CityInfo struct {
//	ActivityID    int64  `gorm:"primary_key;auto_increment" json:"activity_id"`
//	ActivityName  string `gorm:"size(255);not null;" json:"activity_name"`
//	CityId        int64  `gorm:"" json:"city_id"`
//	CityName      string `gorm:"size(255);not null;" json:"city_name"`
//	ImageLocation string `gorm:"size(255);not null;" json:"image_location"`
//	CountryId     int64  `gorm:"" json:"country_id"`
//	CountryName   string `gorm:"size(255);not null;" json:"country_name"`
//	ActivityType  string `gorm:"size(255);not null;" json:"activity_type"`
//	Sponsored     bool   `gorm:"" json:"sponsored"`
//	Points        int64  `gorm:"" json:"points"`
//}

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
	err := db.Debug().DropTableIfExists(&Account{}, &City{}, &Activity{}, &Country{}, &CountryCityActivityMap{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&Account{}, &City{}, &Activity{}, &Country{}, &CountryCityActivityMap{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range Countries {
		err = db.Debug().Model(&Country{}).Create(&Countries[i]).Error
		if err != nil {
			log.Fatalf("cannot seed cards table: %v", err)
		}

	}
	for i, _ := range Cities {
		err = db.Debug().Model(&City{}).Create(&Cities[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range CityInfos {
		err = db.Debug().Model(&Activity{}).Create(&CityInfos[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}

func Load(db *gorm.DB) {
	log.Println("Seed")
	//err := db.Debug().DropTableIfExists(&models.Card{}, &models.User{}).Error
	//if err != nil {
	//	log.Fatalf("cannot drop table: %v", err)
	//}
	err := db.Debug().AutoMigrate(&Account{}, &City{}, &Activity{}, &Country{}, &CountryCityActivityMap{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*
		err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "cards(id)", "cascade", "cascade").Error
		if err != nil {
			log.Fatalf("attaching foreign key error: %v", err)
		}
	*/

	//for i, _ := range card_list {
	//	err = db.Debug().Model(&models.Card{}).Create(&card_list[i]).Error
	//	if err != nil {
	//		log.Fatalf("cannot seed cards table: %v", err)
	//	}
	//
	//}
	//for i, _ := range users {
	//	err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
	//	if err != nil {
	//		log.Fatalf("cannot seed users table: %v", err)
	//	}
	//
	//}
}
