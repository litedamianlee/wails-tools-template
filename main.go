package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"log"
	"runtime"
	"wails-tools-template/backend/app"
	"wails-tools-template/backend/app/constant"
	"wails-tools-template/backend/app/logx"
	"wails-tools-template/backend/service"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	run()
}

func run() {
	opts := &options.App{
		Title:     constant.AppTitle,
		Width:     819,
		Height:    614,
		MinWidth:  819,
		MinHeight: 614,
		// MaxWidth:          1280,
		// MaxHeight:         800,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         runtime.GOOS != "darwin",
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		Menu:             nil,
		Logger:           nil,
		LogLevel:         logger.WARNING,
		OnStartup:        startup,
		OnDomReady:       domReady,
		OnBeforeClose:    beforeClose,
		OnShutdown:       shutdown,
		WindowStartState: options.Normal,
		Bind:             service.Services(),
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    true,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
			BackdropType:        windows.Acrylic,
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHidden(),
			Appearance:           mac.NSAppearanceNameVibrantLight,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   constant.AppTitle,
				Message: "",
				Icon:    icon,
			},
		},
	}
	// Create application with options
	err := wails.Run(opts)
	if err != nil {
		log.Fatal(err)
	}
}

// startup is called at application startup
func startup(ctx context.Context) {
	app.Init()
	logx.Wails().Info("应用已启动")
}

// domReady is called after the front-end dom has been loaded
func domReady(ctx context.Context) {
	logx.Wails().Info("DOM 就绪")
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
func beforeClose(ctx context.Context) (prevent bool) {
	logx.Wails().Info("APP BEFORE CLOSE")
	return false
}

// shutdown is called at application termination
func shutdown(ctx context.Context) {
	logx.Wails().Info("APP SHUTDOWN")
}

// suspend is called when Windows enters low power mode
func suspend() {
	logx.Wails().Info("APP SUSPEND")
}

// resume is called when Windows resumes from low power mode
func resume() {
	logx.Wails().Info("APP RESUME")
}
