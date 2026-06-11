package tag

import (
	"time"
)

type TagModel struct {
	ID uint64 `gorm:"primaryKey"`

	Name string `gorm:"size:100;unique"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (TagModel) TableName() string {
	return "tags"
}
