package deepl

type TranslationResponse struct {
	DetectedSourceLanguage string
	Text string
}

type TranslationRequest struct {
	Text []string
	TargetLang string
}
