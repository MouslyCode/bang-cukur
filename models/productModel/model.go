package productModel

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Price     int64     `gorm:"not null" json:"price"`
	Stock     int64     `gorm:"not null" json:"stock"`
	Img       string    `gorm:"type:varchar(255)" json:"img_url"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
