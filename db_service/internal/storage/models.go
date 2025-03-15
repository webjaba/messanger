package storage

import "time"

type User struct {
	ID               uint32    `gorm:"primaryKey;autoIncrement"`
	Username         string    `gorm:"size:10;unique;not null"`
	Password         string    `gorm:"size:10;not null"`
	MessagesSended   []Message `gorm:"foreignKey:FromUser;references:ID"`
	MessagesRecieved []Message `gorm:"foreignKey:ToUser;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Message struct {
	ID               uint32    `gorm:"primaryKey;autoIncrement"`
	Text             string    `gorm:"type:text;not null"`
	CreatingDateTime time.Time `form:"not null"`
	FromUser         uint32
	ToUser           uint32
}
