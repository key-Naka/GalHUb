package company

import (
	"time"
)

type CompanyModel struct {
	ID uint64 `gorm:"primaryKey"`

	Name string `gorm:"size:100;unique"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (CompanyModel) TableName() string {
	return "companies"
}
