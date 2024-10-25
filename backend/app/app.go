package app

import (
	"context"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
	"wails-tools-template/backend/app/logx"
	"wails-tools-template/backend/storage"
)

var _app = &Application{}

// Application struct
type Application struct {
	once            sync.Once
	log             *logx.Log // loggers for whole application
	storage         *storage.Storage
	currentLanguage string
	ctx             context.Context // wails context
	cache           *cache.Cache
}

func Init() {
	_app.once.Do(func() {
		// setup log
		_app.log = logx.Init()
		_app.storage = storage.NewStorage("XiaLou.db")
		_app.cache = cache.New(15*time.Minute, 10*time.Minute)
	})

}

// Lang get current language
func (a *Application) Lang() string {
	return a.currentLanguage
}

// Ctx get wails context
func (a *Application) Ctx() context.Context {
	return a.ctx
}

// SetCtx set wails context
func (a *Application) SetCtx(ctx context.Context) *Application {
	a.ctx = ctx
	return a
}

// Log get loggers for application
func (a *Application) Log() *logx.Log {
	return a.log
}

func (a *Application) Storage() *storage.Storage {
	return a.storage
}

func (a *Application) Cache() *cache.Cache {
	return a.cache
}
