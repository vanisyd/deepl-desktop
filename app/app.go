package app

import (
	"deepl-desktop/clipboard"
	"deepl-desktop/deepl"
	"sync"
)

var Clipboard clipboard.Clipboard
var mu sync.Mutex

func Init() {
	err := Clipboard.Init()
	if err != nil {
		panic(err)
	}
}

func Handle() {
	mu.Lock()
	text, err := Clipboard.Read()
	if err != nil {
		panic(err)
	}
	deepl.TranslateText(text)
	mu.Unlock()
}
