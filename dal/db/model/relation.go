package model

import "gorm.io/gorm"

type Relation struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;uniqueIndex:idx_relation"`
	ToUserId int64 `gorm:"column:to_user_id;not null;uniqueIndex:idx_relation"`
	Action   bool  `gorm:"column:action" json:"action"`
}

func (Relation) TableName() string {
	return "relation"
}
