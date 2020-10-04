package model

import (
	"time"

	"gorm.io/gorm"
)

type GORMBase struct {
	ID        int   `json:"id" gorm:"AUTO_INCREMENT"`
	CreatedAt int64 `json:"createAt"`
	UpdatedAt int64 `json:"updateAt"`
}

func (m *GORMBase) BeforeCreate(scope *gorm.DB) error {
	if m.UpdatedAt == 0 {
		scope.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	}

	scope.Statement.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}

func (m *GORMBase) BeforeUpdate(scope *gorm.DB) error {
	scope.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}
