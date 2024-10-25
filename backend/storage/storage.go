package storage

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"wails-tools-template/backend/storage/gen"
	"wails-tools-template/backend/utils"
)

var (
	ErrMustProvideDir  = errors.New("must provide config dir")
	ErrMustContainDB   = errors.New("must contain .db")
	ErrQueryEmpty      = errors.New("no query provided")
	ErrFailedConnectDB = errors.New("failed to connect to database")

	LogPrefixStorage = "STORAGE"
)

type Storage struct {
	db  *gorm.DB
	orm *gen.Query
	log *utils.Logger
}

func NewStorage(filename string) *Storage {
	err := Setup()
	if err != nil {
		panic(err)
		return nil
	}
	d := &Storage{}
	dbPath := filepath.Join(Directory(), filename)
	log.Println(fmt.Sprintf("db path: %s", dbPath))
	db, err := NewSqlite(dbPath, nil)
	if err != nil {
		panic(db)
	}
	d.db = db
	d.orm = gen.Use(db)
	d.log = NewLog()
	return d
}

func NewLog() *utils.Logger {
	return utils.NewConsoleLogger(LogPrefixStorage)

}

func Setup() error {
	dbPath := Directory()
	err := os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Directory() string {
	configDir, _ := os.UserConfigDir()
	return filepath.Join(configDir, "XiuXianHelpTool")
}
