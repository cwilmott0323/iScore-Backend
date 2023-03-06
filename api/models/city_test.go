package models

//
//import (
//	"errors"
//	"fmt"
//	"gorm.io/gorm"
//	"github.com/stretchr/testify/assert"
//	"testing"
//
//	_ "github.com/lib/pq" // postgres driver
//	"github.com/orlangure/gnomock"
//	"github.com/orlangure/gnomock/preset/postgres"
//)
//
//func TestGetCities(t *testing.T) {
//	c := City{}
//
//	db := Prep(t)
//
//	// Seed data
//	TestData(db)
//
//	d, err := c.GetCities(db, "England")
//
//	if err != nil {
//		return
//	}
//	assert.Equal(t, d[0].CityName, "London", "The two words should be the same.")
//	assert.Equal(t, len(d), 1, "One object should be returned")
//}
//
//func TestGetCitiesError(t *testing.T) {
//	c := City{}
//
//	db := Prep(t)
//	// Seed data
//	TestData(db)
//	expectedError := errors.New("no Cities Returned")
//
//	_, err := c.GetCities(db, "not Found City")
//	if assert.Error(t, err) {
//		assert.Equal(t, expectedError, err)
//	}
//}
//
//func TestGetCitiesInfo(t *testing.T) {
//	c := CityInfo{}
//
//	db := Prep(t)
//
//	// Seed data
//	TestData(db)
//
//	d, err := c.GetCitiesInfo(db, "England", "London")
//
//	if err != nil {
//		return
//	}
//	assert.Equal(t, d[0].ActivityName, "Big Ben", "The two words should be the same.")
//	assert.Equal(t, len(d), 1, "One object should be returned")
//}
//
//func TestGetCitiesInfoError(t *testing.T) {
//	c := CityInfo{}
//
//	db := Prep(t)
//
//	// Seed data
//	TestData(db)
//
//	expectedError := errors.New("no Cities Returned")
//	_, err := c.GetCitiesInfo(db, "Fake Country", "Fake City")
//	if assert.Error(t, err) {
//		assert.Equal(t, expectedError, err)
//	}
//}
//
//func Prep(t *testing.T) *gorm.DB {
//	p := postgres.Preset(
//		postgres.WithUser("gnomock", "gnomick"),
//		postgres.WithDatabase("api"),
//	)
//	container, _ := gnomock.Start(p)
//	t.Cleanup(func() { _ = gnomock.Stop(container) })
//
//	connStr := fmt.Sprintf(
//		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
//		container.Host, container.DefaultPort(),
//		"gnomock", "gnomick", "api",
//	)
//	db, _ := gorm.Open("postgres", connStr)
//	return db
//}
