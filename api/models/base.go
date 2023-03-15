package models

type CountryCityActivityMap struct {
	//gorm.Model
	ActivityID int64    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;unique" json:"activity_id"`
	CityID     int64    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"city_id"`
	CountryID  int64    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"country_id"`
	Activity   Activity `gorm:"foreignKey:ActivityID"`
	City       City     `gorm:"foreignKey:CityID"`
	Country    Country  `gorm:"foreignKey:CountryID"`
}
