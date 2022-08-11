package interactor

import "github.com/shumon84/minecraft-discord-bot/pkg/core"

type MockLangUsecase struct {
	DeathMessageToLangKeyAndArgsExpectBehavior func(message string, langCode core.LangCode) (string, []string, error)
	EntityToLangKeyAndArgsExpectBehavior       func(message string, langCode core.LangCode) (string, []string, error)
	GetFixedTextExpectBehavior                 func(langCode core.LangCode, key string, args []string) (string, error)
}

func (m *MockLangUsecase) DeathMessageToLangKeyAndArgs(message string, langCode core.LangCode) (string, []string, error) {
	return m.DeathMessageToLangKeyAndArgsExpectBehavior(message, langCode)
}
func (m *MockLangUsecase) EntityToLangKeyAndArgs(message string, langCode core.LangCode) (string, []string, error) {
	return m.EntityToLangKeyAndArgsExpectBehavior(message, langCode)
}
func (m *MockLangUsecase) GetFixedText(langCode core.LangCode, key string, args []string) (string, error) {
	return m.GetFixedTextExpectBehavior(langCode, key, args)
}
