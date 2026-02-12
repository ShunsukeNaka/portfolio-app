package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username         string    `gorm:"size:50;not null"`
	Email            string    `gorm:"size:255;unique;not null"`
	XAccount         string    `gorm:"size:50"`
	InstagramAccount string    `gorm:"size:50"`
	ProfileText      string    `gorm:"size:200"`
	OwnedPets        string    `gorm:"size:100"`
	AvatarURL        string    `gorm:"type:text"`
	CreatedAt        time.Time
	Articles         []Article `gorm:"constraint:OnDelete:CASCADE;"`
	Likes            []Article `gorm:"many2many:likes;"`
	Comments         []Comment `gorm:"constraint:OnDelete:CASCADE;"`
}

type Article struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Title     string    `gorm:"size:255;not null"`
	Content   string    `gorm:"type:text;not null"`
	PetType   string    `gorm:"size:50;not null;index"`
	PetSize   string    `gorm:"size:20;not null;index"`
	CreatedAt time.Time
	Tags      []Tag `gorm:"many2many:article_tags;constraint:OnDelete:CASCADE;"`
}

type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:50;unique;not null"`
}

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ArticleID uuid.UUID `gorm:"type:uuid;not null"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time
}

type Follow struct {
	FollowerID  uuid.UUID `gorm:"primaryKey"`
	FollowingID uuid.UUID `gorm:"primaryKey"`
	CreatedAt   time.Time
}
