package deepl

import (
	"fmt"
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
		panic(err)
	}

	if len(response.Translations) > 0 {
		fmt.Printf("Source text: %s | Source lang: %s | Translation: %s\n", text, response.Translations[0].DetectedSourceLanguage, response.Translations[0].Text)
	}
}
