package config

import (
	"github.com/wailsapp/wails"
	"os"
	"path/filepath"
)

type Config struct {
	Runtime *wails.Runtime
	Logger  *wails.CustomLogger
}

func NewConfig() *Config {
	return &Config{}
}

func Directory() string {
	configDir, _ := os.UserConfigDir()
	return filepath.Join(configDir, wails.AppTitleEn)
}

func Setup() error {
	dbPath := Directory()
	err := os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
