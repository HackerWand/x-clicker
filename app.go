package main

import (
	"context"
	"github.com/go-vgo/robotgo"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	_ "golang.design/x/hotkey/mainthread"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

type Config struct {
	Type     string
	Double   bool
	Speed    int
	StartKey string
	StopKey  string
	Running  bool
}

var config = Config{
	Type:     "left",
	Double:   false,
	Speed:    1000,
	StartKey: "F1",
	StopKey:  "F2",
	Running:  false,
}

var hotkeys *HotKeys

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	hotkeys = NewHotKeys(a)
	a.SetMenu()
}

func (a *App) beforeClose(ctx context.Context) bool {
	hotkeys.UnBind()
	return true
}

func (a *App) SetMenu() {
	hotkeys.Bind(config.StartKey, config.StopKey)
	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("x-clicker")
	FileMenu.AddText("Start", keys.Key(config.StartKey), func(data *menu.CallbackData) {
		a.Start()
	})
	FileMenu.AddText("Stop", keys.Key(config.StopKey), func(data *menu.CallbackData) {
		a.Stop()
	})
	runtime.MenuSetApplicationMenu(a.ctx, AppMenu)
	runtime.MenuUpdateApplicationMenu(a.ctx)
}

func (a *App) SetConfig(newConfig Config) {
	config = newConfig
	a.SetMenu()
}

func (a *App) GetConfig() Config {
	return config
}

func (a *App) Start() {
	if config.Running == true {
		return
	}
	config.Running = true
	for {
		if !config.Running {
			break
		}
		robotgo.Click(config.Type, config.Double)
		time.Sleep(time.Duration(config.Speed) * time.Millisecond)
	}
}

func (a *App) Stop() {
	config.Running = false
}
