package entity

import "gorm.io/gorm"

// History
// @Description:历史消息实体
type History struct {
	gorm.Model
	Exp string
	Val string
}
