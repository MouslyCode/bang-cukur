package itemModel

import (
	"time"

	"github.com/MouslyCode/bang-cukur/common/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID        uuid.UUID         `gorm:"type:char(36);primaryKey;not null" json:"id"`
	Name      string            `gorm:"type:varchar(100);not null" json:"name"`
	Price     int64             `gorm:"not null" json:"price"`
	Stock     *int              `gorm:"type:int" json:"stock,omitempty"`
	Img       string            `gorm:"type:varchar(255)" json:"img_url"`
	Type      constant.ItemType `gorm:"type:varchar(20);not null" json:"type"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	return
}
