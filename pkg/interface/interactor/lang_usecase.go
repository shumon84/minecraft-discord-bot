package interactor

import (
	"errors"

	"github.com/shumon84/minecraft-discord-bot/pkg/core"
)

var ErrLangUseCaseNotFoundMatchedPlaceHolder = errors.New("not found matched place holder")

type LangUsecase interface {
	DeathMessageToLangKeyAndArgs(message string, langCode core.LangCode) (string, []string, error)
	EntityToLangKeyAndArgs(message string, langCode core.LangCode) (string, []string, error)
	GetFixedText(langCode core.LangCode, key string, args []string) (string, error)
}
