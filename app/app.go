package app

import "deepl-desktop/deepl"

func Handle() {
	//text := string(clipboard.Read(clipboard.FmtText))
	deepl.TranslateText("My little friend")
}
