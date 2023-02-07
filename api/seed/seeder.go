package seed

import (
	"github.com/jinzhu/gorm"
	"iScore-api/api/models"
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

func Load(db *gorm.DB) {
	log.Println("Seed")
	//err := db.Debug().DropTableIfExists(&models.Card{}, &models.User{}).Error
	//if err != nil {
	//	log.Fatalf("cannot drop table: %v", err)
	//}
	err := db.Debug().AutoMigrate(&models.Account{}).Error
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
