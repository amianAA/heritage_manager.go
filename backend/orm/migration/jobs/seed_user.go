package jobs

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"heritago/backend/orm/models"
	"strings"
)

var (
	fname    = "Alvaro"
	lname    = "Amian"
	country  = "Spain"
	password = models.HashAndSaltPwd("0987654321")
	userID   = strings.ToLower(fname[:1] + lname)

	firstUser = &models.User{
		Email:     "amianalvaro@gmail.com",
		UserID:    &userID,
		FirstName: &fname,
		LastName:  &lname,
		Country:   &country,
		Password:  &password,
	}
)

// SeedUsers inserts the first users
var SeedUsers = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstUser).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstUser).Error
	},
}