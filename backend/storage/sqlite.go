package storage

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
	"strings"
)

func NewSqlite(path string, migrations ...interface{}) (*gorm.DB, error) {
	log.Println(fmt.Sprintf("path: %s", path))
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, ErrFailedConnectDB
	}
	//err = db.AutoMigrate(migrations...)
	return db, err
}

func HasKeywords(key, value string) bool {
	return strings.Contains(value, key+":")
}

func ParseKeywords(key string, value string) string {
	if !HasKeywords(key, value) {
		return key
	}
	split := strings.SplitAfter(value, ":")
	q := strings.TrimSpace(split[1])
	return q
}
