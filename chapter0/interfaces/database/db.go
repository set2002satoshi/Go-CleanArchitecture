package database

import (
	"gorm.io/gorm"
)

type DB interface {
	Begin() *gorm.DB
	Connect() *gorm.DB
}
