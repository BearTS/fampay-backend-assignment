package db

import (
	"time"

	"gorm.io/gorm"
)

type Videos struct {
	gorm.Model

	VideoId      string    `gorm:"column:video_id;uniqueIndex;not null;type:text"`
	Title        string    `gorm:"column:title;not null;type:text"`
	Description  string    `gorm:"column:description;type:text"`
	Thumbnail    string    `gorm:"column:thumbnail;type:text"`
	ChannelTitle string    `gorm:"column:channel_title;type:text"`
	PublishedAt  time.Time `gorm:"column:published_at;not null;type:timestamp"`
}

func (Videos) TableName() string {
	return "videos"
}

func (db *db) CreateVideosBulk(videos []*Videos) error {
	return db.gormDB.Create(&videos).Error
}
