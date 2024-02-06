package pkg

import (
	"github.com/BearTS/fampay-backend-assignment/pkg/db"
	"gorm.io/gorm"
)

type Migrate struct {
	TableName string
	Run       func(*gorm.DB) error
}

func AutoMigrate(gormD *gorm.DB) []Migrate {
	var videos db.Videos

	videosM := Migrate{TableName: videos.TableName(),
		Run: func(gormD *gorm.DB) error { return gormD.AutoMigrate(&videos) }}

	return []Migrate{
		videosM,
	}
}
