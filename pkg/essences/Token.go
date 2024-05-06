package essences

import "time"

type Token struct {
	Plaintext string    `json:"token"`
	Hash      []byte    `gorm:"type:bytea;primaryKey"`
	UserID    uint      `gorm:"not null"`
	Expiry    time.Time `gorm:"not null"`
	Scope     string    `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
}
