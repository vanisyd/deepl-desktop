package main

import (
	"deepl-desktop/app"
	"github.com/joho/godotenv"
	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
	"log"
	"sync"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	mainthread.Init(fn)
}

func fn() {
	if err := clipboard.Init(); err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			keyHandler()
		}
	}()
	wg.Wait()
}

func keyHandler() {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModShift}, hotkey.KeyC)
	if err := hk.Register(); err != nil {
		return
	}

	<-hk.Keydown()
	app.Handle()

	hk.Unregister()
	return
}
