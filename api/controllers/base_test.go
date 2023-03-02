package controllers

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

func TestOpenDB(t *testing.T) {
	mockError := errors.New("uh oh")
	subtests := []struct {
		name                 string
		d, u, p, a, db, port string
		sqlOpener            func(dialect string, args ...interface{}) (db *gorm.DB, err error)
		expectedErr          error
	}{
		{
			name: "happy path",
			d:    "postgres",
			u:    "u",
			p:    "p",
			a:    "a",
			db:   "db",
			port: "4000",
			sqlOpener: func(s string, s2 ...interface{}) (db *gorm.DB, err error) {
				fmt.Println(s, s2)
				if s != "postgres" {
					return nil, errors.New("wrong connection string")
				}
				return nil, nil
			},
		},
		{
			name: "error from sqlOpener",
			sqlOpener: func(s string, s2 ...interface{}) (db *gorm.DB, err error) {
				return nil, mockError
			},
			expectedErr: mockError,
		},
	}
	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			_, err := OpenDB(subtest.d, subtest.u, subtest.p, subtest.port, subtest.a, subtest.db, subtest.sqlOpener)
			if !errors.Is(err, subtest.expectedErr) {
				t.Errorf("expected error (%v), got error (%v)", subtest.expectedErr, err)
			}
		})
	}
}
