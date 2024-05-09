package app

import (
	"github.com/gen2brain/beeep"
	"golang.design/x/clipboard"
)

func Handle() {
	text := string(clipboard.Read(clipboard.FmtText))
	beeep.Notify("DeepL Desktop", text, "")
}
