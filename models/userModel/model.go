package userModel

import (
	"time"

	"github.com/MouslyCode/bang-cukur/common/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(120);not null" json:"-"`
	RoleID    uint      `json:"-"`
	Role      Role      `gorm:"foreignKey:RoleID" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate is a hook function that is called before the User is created.
// It is used to generate a UUID for the User and hash the password.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	hashed, err := helper.HashPassword(u.Password)
	if err != nil {
		return
	}

	u.Password = hashed
	return
}

type Role struct {
	ID        uint   `gorm:"primaryKey" json:"-"`
	Name      string `gorm:"type:varchar(20);uniqueIndex;not null" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
