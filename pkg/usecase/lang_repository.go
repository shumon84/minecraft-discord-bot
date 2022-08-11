package usecase

import (
	"github.com/shumon84/minecraft-discord-bot/pkg/core"
	"github.com/shumon84/minecraft-discord-bot/pkg/domain/entity"
)

type LangRepository interface {
	Find(langCode core.LangCode) (*entity.Lang, error)
}
