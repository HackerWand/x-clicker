package main

import (
	"golang.design/x/hotkey"
	"sync"
)

var CharKeys = map[string]uint8{
	"F1":  0x7A,
	"F2":  0x78,
	"F3":  0x63,
	"F4":  0x76,
	"F5":  0x60,
	"F6":  0x61,
	"F7":  0x62,
	"F8":  0x64,
	"F9":  0x65,
	"F10": 0x6D,
	"F11": 0x67,
	"F12": 0x6F,
}

type HotKeys struct {
	wg          *sync.WaitGroup
	startHotKey *hotkey.Hotkey
	endHotKey   *hotkey.Hotkey
	startKey    hotkey.Key
	endKey      hotkey.Key
	close       chan bool
	app         *App
}

func NewHotKeys(app *App) *HotKeys {
	return &HotKeys{
		app: app,
	}
}

func (hot *HotKeys) Bind(start string, end string) {
	hot.startKey = hotkey.Key(CharKeys[start])
	hot.endKey = hotkey.Key(CharKeys[end])

	if hot.wg != nil {
		hot.UnBind()
	}

	//go mainthread.Init(hot.init)
	hot.init()
}

func (hot *HotKeys) init() {
	hot.close = make(chan bool)
	hot.wg = &sync.WaitGroup{}
	hot.wg.Add(2)

	hot.startHotKey = hotkey.New([]hotkey.Modifier{}, hot.startKey)
	if err := hot.startHotKey.Register(); err != nil {
		panic(err)
	}
	go func() {
		defer hot.wg.Done()
		for {
			select {
			case <-hot.close:
				if err := hot.startHotKey.Unregister(); err != nil {
					panic(err)
				}
				return
			case <-hot.startHotKey.Keydown():
				hot.app.Start()
			}
		}
	}()

	hot.endHotKey = hotkey.New([]hotkey.Modifier{}, hot.endKey)
	if err := hot.endHotKey.Register(); err != nil {
		panic(err)
	}
	go func() {
		defer hot.wg.Done()
		for {
			select {
			case <-hot.close:
				if err := hot.endHotKey.Unregister(); err != nil {
					panic(err)
				}
				return
			case <-hot.endHotKey.Keydown():
				hot.app.Stop()
			}
		}
	}()

	hot.wg.Wait()
}

func (hot *HotKeys) UnBind() {
	close(hot.close)
	hot.wg.Wait()
}
