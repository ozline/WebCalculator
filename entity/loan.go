package entity

import (
	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	Duration float64
	Rate     float64
}
