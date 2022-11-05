package entity

import (
	"errors"
	"time"

	"github.com/set2002satoshi/Go-CleanArchitecture/chapter1/pkg/module/temporary"
)

type ActiveUserEntity struct {
	ActiveUserID        temporary.IDENTIFICATION `gorm:"primaryKey"`
	ActiveUserDetailsID uint
	NickName            string             `gorm:"size:16;not null"`
	Bio                 string             `gorm:"size:256;not null"`
	ActiveBlog          []ActiveBlogEntity `gorm:"foreignKey:ActiveID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Revision            temporary.REVISION `gorm:"not null"`
}

func NewActiveUserEntity(
	id int,
	nickName,
	bio string,
	blog []ActiveBlogEntity,
	createdAt,
	updatedAt time.Time,
	revision temporary.REVISION,
) (*ActiveUserEntity, error) {
	u := &ActiveUserEntity{}

	if u.setNickname(nickName) {
		return nil, errors.New("Nameセッターのエラーが出てるよ")
	}

	return u, nil
}

func (u *ActiveUserEntity) setNickname(nickName string) bool {
	u.NickName = nickName
	return false
}
