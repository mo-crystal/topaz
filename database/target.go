package database

import (
	"errors"

	"gorm.io/gorm"
)

type Target struct {
	Id          int
	CallbackUrl string
	Enabled     bool
}

func SaveTarget(t *Target) {
	db.Save(t)
}

func FindTarget(callbackUrl string) *Target {
	target := Target{}
	result := db.Where("callback_url = ? and enabled = 1", callbackUrl).First(&target)
	if errors.Is(gorm.ErrRecordNotFound, result.Error) {
		return nil
	} else {
		return &target
	}
}
