package storage

type User struct {
	ID               uint      `gorm:"primaryKey"`
	Username         string    `gorm:"size:10;unique;not null"`
	Password         string    `gorm:"size:10;not null"`
	MessagesSended   []Message `gorm:"foreignKey:FromUser;references:ID"`
	MessagesRecieved []Message `gorm:"foreignKey:ToUser;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Message struct {
	ID           uint   `gorm:"primaryKey"`
	Text         string `gorm:"type:text;not null"`
	CreatingTime string `gorm:"type:time;not null"`
	CreatingDate string `form:"type:date;not null"`
	FromUser     uint
	ToUser       uint
}

// db.Model(&Message{}).AddForeignKey("from_user", "users(id)", "CASCADE", "CASCADE")
