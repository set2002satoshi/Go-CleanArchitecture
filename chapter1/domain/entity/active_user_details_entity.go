package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/set2002satoshi/Go-CleanArchitecture/chapter1/pkg/module/temporary"
)

type ActiveUserDetailsEntity struct {
	ActiveUserDetailsID temporary.IDENTIFICATION `gorm:"primaryKey"`
	Email               string                   `gorm:"unique;not null"`
	Password            []byte                   `gorm:"not null"`
	User                ActiveUserEntity         `gorm:"foreignKey:ActiveUserDetailsID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Revision            temporary.REVISION `gorm:"not null"`
}

func NewActiveUserDetailsEntity(
	id int,
	name,
	email,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
	revision temporary.REVISION,
) (*ActiveUserDetailsEntity, error) {
	u := &ActiveUserDetailsEntity{}

	if u.setID(id) {
		return nil, errors.New("idセッターのエラーが出てるよ")
	}

	if u.setEmail(email) {
		return nil, errors.New("ScreenNameセッターのエラーが出てるよ")
	}
	if u.setPassword(password) {
		return nil, errors.New("ScreenNameセッターのエラーが出てるよ")
	}
	if u.setCreatedAt(createdAt) {
		return nil, errors.New("setCreatedAtにエラーが発生しています。")
	}
	if u.setUpdatedAt(updatedAt) {
		return nil, errors.New("setUpdatedAtにエラーが発生しています。")
	}

	if u.setRevision(revision) {
		return nil, errors.New("setRevisionにエラーが発生しています。")
	}
	return u, nil
}

func (u *ActiveUserDetailsEntity) setID(id int) bool {
	i, err := temporary.NewIDENTIFICATION(id)
	if err != nil {
		return true
	}
	u.ActiveUserDetailsID = i
	return false
}

func (u *ActiveUserDetailsEntity) setEmail(email string) bool {
	u.Email = email
	return false
}

func (u *ActiveUserDetailsEntity) setPassword(password string) bool {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return true
	}
	u.Password = []byte(pass)
	return false
	// u.Password = []byte(password)
}

func (s *ActiveUserDetailsEntity) setCreatedAt(createdAt time.Time) bool {
	s.CreatedAt = createdAt
	return false
}

func (s *ActiveUserDetailsEntity) setUpdatedAt(updatedAt time.Time) bool {
	s.CreatedAt = updatedAt
	return false
}

func (s *ActiveUserDetailsEntity) setRevision(revision temporary.REVISION) bool {
	s.Revision = revision
	return false
}

// func (s *ActiveUserDetailsEntity) CountUpRevisionNumber(num temporary.REVISION) error {

// 	if s.GetRevision() != num {
// 		return errors.New("改定番号が異なるため更新はできません")
// 	}
// 	if ok := s.setRevision(num + 1); ok {
// 		return errors.New("Invalid setting")
// 	}
// 	return nil
// }
