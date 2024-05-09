package deepl

import (
	"fmt"
	"github.com/gen2brain/beeep"
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

	response := TranslationsResponse{}
	err := httpClient.SendRequest(request, &response)
	if err != nil {
		panic(fmt.Sprintf("Translation error: %v", err))
	}

	if len(response.Translations) > 0 {
		beeep.Notify(text, response.Translations[0].Text, "")
	}
}
