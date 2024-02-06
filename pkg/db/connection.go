/*
 * File: connection.go
 * Author: Anuj Parihar <anujparihar@yahoo.com>
 * GitHub: github.com/bearts
 * Copyright © 2024 by Anuj Parihar
 */

package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/BearTS/fampay-backend-assignment/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, *sql.DB) {

	host := config.Config.PostgresHost
	user := config.Config.PostgresUser
	password := config.Config.PostgresPass
	dbName := config.Config.PostgresDb
	port := fmt.Sprintf("%d", config.Config.PostgresPort)

	dsn := "host=" + host +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbName +
		" port=" + port +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[Connection], Error in opening db")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("[Connection], Error in setting sqldb")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, sqlDB
}
