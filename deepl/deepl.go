package deepl

import (
	"github.com/gen2brain/beeep"
	"log"
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
		log.Fatalf("Translation error: %v", err)
	}

	if len(response.Translations) > 0 {
		beeep.Notify(text, response.Translations[0].Text, "")
	}
}
