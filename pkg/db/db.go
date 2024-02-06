package db

import "gorm.io/gorm"

type db struct {
	gormDB *gorm.DB
}

type DB interface {
	BeginTxx() *gorm.DB
	CommitTxx(tx *gorm.DB) error
}

// NewDB creates a new DB object.
// It takes a *gorm.DB object representing the database connection.
// Returns a DB object.
func NewDB(gormDB *gorm.DB) DB {
	return &db{gormDB: gormDB}
}

func (db *db) BeginTxx() *gorm.DB {
	return db.gormDB.Begin()
}

// CommitTxx commits the transaction and returns any error that occurred.
// It takes a *gorm.DB object representing the transaction to be committed.
// Returns an error if the commit operation fails.
func (db *db) CommitTxx(tx *gorm.DB) error {
	return tx.Commit().Error
}
