package usecase

import (
	"github.com/shumon84/minecraft-discord-bot/pkg/core"
	"github.com/shumon84/minecraft-discord-bot/pkg/domain/entity"
)

type MockLangRepository struct {
	FindExpectBehavior func(core.LangCode) (*entity.Lang, error)
}

var _ = LangRepository(new(MockLangRepository))

func (m *MockLangRepository) Find(langCode core.LangCode) (*entity.Lang, error) {
	return m.FindExpectBehavior(langCode)
}
