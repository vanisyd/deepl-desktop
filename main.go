package main

import (
	"deepl-desktop/deepl"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	deepl.TranslateText("My first test")
}

/*func fn() {
	if err := clipboard.Init(); err != nil {
		log.Fatalf("Error initializing clipboard: %v", err)
	}

	wg.Add(1)
	go func() {
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
	defer hk.Unregister()

	log.Printf("Registered hotkey: %v", hk)
	<-hk.Keydown()
	log.Printf("Hotkey pressed: %v", hk)
	app.Handle()

	return
}*/
