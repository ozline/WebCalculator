package entity

import (
	"gorm.io/gorm"
)

type Deposit struct {
	gorm.Model
	Duration float64
	Rate     float64
}
