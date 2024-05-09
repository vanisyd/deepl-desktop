package deepl

type TranslationsResponse struct {
	Translations []TranslationResponse `json:"translations"`
}

type TranslationResponse struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}

type TranslationRequest struct {
	Text       []string `json:"text"`
	TargetLang string   `json:"target_lang"`
}
