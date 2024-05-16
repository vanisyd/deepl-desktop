package main

import (
	"deepl-desktop/app"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vanisyd/hotkey"
	"github.com/vanisyd/hotkey/input"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	app.Init()

	hk := hotkey.Hotkey{
		Keys: []input.KeyCode{
			input.KeyCtrl,
			input.KeyC,
		},
		TapsCount: 2,
	}
	go hk.Register()

	for {
		select {
		case <-hk.HotkeyPressed:
			fmt.Println("Hotkey pressed")
			app.Handle()
		default:
		}
	}
}
