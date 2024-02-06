package db

import (
	"time"

	"gorm.io/gorm"
)

type Videos struct {
	gorm.DB

	Title        string    `gorm:"column:title;not null;type:text"`
	Description  string    `gorm:"column:description;not null;type:text"`
	Thumbnail    string    `gorm:"column:thumbnail;not null;type:text"`
	ChannelTitle string    `gorm:"column:channel_title;not null;type:text"`
	PublishedAt  time.Time `gorm:"column:published_at;not null;type:timestamp"`
}

func (Videos) TableName() string {
	return "videos"
}

