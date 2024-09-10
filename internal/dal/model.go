package dal

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Password string `gorm:"type:varchar(255);not null"`
}

type PersonalInfo struct {
	UserID    uint   `gorm:"primaryKey"`
	UserName  string `gorm:"type:varchar(255);not null"`
	Gender    int    `gorm:"check:gender >= 0 AND gender <= 2"`
	Birthday  time.Time
	Signature string `gorm:"type:text"`
	Email     string `gorm:"type:varchar(255);unique"`
	Avatar    []byte `gorm:"type:bytea"`
	User      User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Competence struct {
	UserID     uint `gorm:"primaryKey"`
	Competence int  `gorm:"primaryKey;check:competence >= 0 AND competence <= 10"`
	User       User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Post struct {
	ID      uint      `gorm:"primaryKey;autoIncrement"`
	UserID  uint      `gorm:"not null"`
	Date    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	Title   string    `gorm:"type:varchar(255);not null"`
	Details string    `gorm:"type:text"`
	User    User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Comment struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	UserID  uint `gorm:"not null"`
	PostID  uint `gorm:"not null"`
	ReplyID uint
	Date    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	Details string    `gorm:"type:text"`
	User    User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Post    Post      `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
	Reply   *Comment  `gorm:"foreignKey:ReplyID;constraint:OnDelete:CASCADE"`
}
