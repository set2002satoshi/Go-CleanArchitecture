package entity

import (
	"time"

	"github.com/set2002satoshi/Go-CleanArchitecture/chapter1/pkg/module/temporary"
)

type ActiveBlogEntity struct {
	ActiveBlogID temporary.IDENTIFICATION `gorm:"primaryKey"`
	ActiveUserID uint
	Title        string `gorm:"size:25;not null"`
	Context      string `gorm:"size:256;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Revision     temporary.REVISION `gorm:"not null"`
}
