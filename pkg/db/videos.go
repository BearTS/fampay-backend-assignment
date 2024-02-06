package db

import (
	"strings"
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
	VideoUrl     string    `gorm:"column:video_url;type:text"`
	PublishedAt  time.Time `gorm:"column:published_at;not null;type:timestamp"`
}

func (Videos) TableName() string {
	return "videos"
}

func (db *db) CreateVideosBulk(videos []*Videos) error {
	return db.gormDB.Create(&videos).Error
}

func (db *db) GetAllVideosPaginated(offset int, limit int, title *string, description *string) ([]*Videos, error) {
	var videos []*Videos
	query := db.gormDB

	// TODO: Better way to do this
	if title != nil {
		arrayOfString := strings.Split(*title, " ")
		for _, word := range arrayOfString {
			query = query.Where("title LIKE ?", "%"+word+"%")
		}
	}
	if description != nil {
		arrayOfString := strings.Split(*description, " ")
		for _, word := range arrayOfString {
			query = query.Where("description LIKE ?", "%"+word+"%")
		}
	}
	err := query.Offset(offset).Limit(limit).Order("published_at DESC").Find(&videos).Error
	return videos, err
}

func (db *db) GetAllVideosCount() (int64, error) {
	var count int64
	err := db.gormDB.Model(&Videos{}).Count(&count).Error
	return count, err
}
