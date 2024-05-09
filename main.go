package main

import (
	"deepl-desktop/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	app.Handle()
	/*if err := clipboard.Init(); err != nil {
		panic(fmt.Sprintf("Error initializing clipboard: %v", err))
	}*/
	//mainthread.Init(fn)
}

/*func fn() {

	wg := sync.WaitGroup{}
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

	log.Printf("Registered hotkey: %v", hk)
	<-hk.Keydown()
	app.Handle()

	hk.Unregister()
	return
}*/
