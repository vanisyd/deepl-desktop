package deepl

import (
	"net/http"
)

func TranslateText(text string) {
	httpClient := HTTPClient{}
	httpClient.NewClient()

	requestBody := TranslationRequest{
		Text:       []string{text},
		TargetLang: langUK,
	}
	request := Request{
		Method:   http.MethodPost,
		Endpoint: APITranslateText,
		Body:     requestBody,
	}

}
