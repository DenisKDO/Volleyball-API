package essences

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Token struct {
	gorm.Model
	Plaintext TokenHash `gorm:"not null"`
	Hash      []byte    `gorm:"type:bytea;primaryKey"`
	UserID    uint      `gorm:"not null"`
	Expiry    time.Time `gorm:"not null"`
	Scope     string    `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
}

type TokenHash struct {
	Ptext string
}
