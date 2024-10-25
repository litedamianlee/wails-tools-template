package logx

import (
	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
	gormLogger "gorm.io/gorm/logger"
	"os"
	"wails-tools-template/backend/utils"
)

var log *Log

const (
	LogPrefixDatabase = "DBS"
	LogPrefixWails    = "WLS"
	LogPrefixI18n     = "I18"
	LogPrefixTray     = "TRY"
	LogPrefixServices = "SEV"
)

type Log struct {
	database gormLogger.Interface
	wails    wailsLogger.Logger
	i18n     *utils.Logger
	tray     *utils.Logger
	services *utils.Logger
}

func Init() *Log {
	log = newConsoleLogger()
	return log
}

func newConsoleLogger() *Log {
	return &Log{
		database: utils.NewGormConsoleLogger(LogPrefixDatabase),
		wails:    utils.NewWailsConsoleLogger(LogPrefixWails),
		i18n:     utils.NewConsoleLogger(LogPrefixI18n),
		tray:     utils.NewConsoleLogger(LogPrefixTray),
		services: utils.NewConsoleLogger(LogPrefixServices),
	}
}

func NewFileLogger(logPath string) *Log {
	logFile, err := os.OpenFile(
		logPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		utils.Utils().Panic("failed to open log file: " + err.Error())
	}
	return &Log{
		database: utils.NewGormFileLogger(LogPrefixDatabase, logFile),
		wails:    utils.NewWailsFileLogger(LogPrefixWails, logFile),
		i18n:     utils.NewFileLogger(LogPrefixI18n, logFile),
		tray:     utils.NewFileLogger(LogPrefixTray, logFile),
		services: utils.NewFileLogger(LogPrefixServices, logFile),
	}
}

func Wails() wailsLogger.Logger {
	return log.wails
}

func Tray() *utils.Logger {
	return log.tray
}

func Services() *utils.Logger {
	return log.services
}

func Database() gormLogger.Interface {
	return log.database
}

func I18n() *utils.Logger {
	return log.i18n
}
