package interactor

import (
	"errors"

	"github.com/shumon84/minecraft-discord-bot/pkg/core"
)

type LangInteractor struct {
	lu LangUsecase
}

func NewLangInteractor(lu LangUsecase) *LangInteractor {
	return &LangInteractor{
		lu: lu,
	}
}

func (li *LangInteractor) TranslateDeathMessage(message string, fromCode core.LangCode, toCode core.LangCode) (string, error) {
	deathMessageKey, args, err := li.lu.DeathMessageToLangKeyAndArgs(message, fromCode)
	if err != nil {
		return "", err
	}

	translatedArgs := make([]string, 0, len(args))
	for _, arg := range args {
		nameKey, nameArgs, err := li.lu.EntityToLangKeyAndArgs(arg, fromCode)
		if errors.Is(err, ErrLangUseCaseNotFoundMatchedPlaceHolder) {
			translatedArgs = append(translatedArgs, arg)
			continue
		}
		if err != nil {
			return "", err
		}
		translatedArg, err := li.lu.GetFixedText(toCode, nameKey, nameArgs)
		if err != nil {
			return "", err
		}
		translatedArgs = append(translatedArgs, translatedArg)
	}

	translatedMessage, err := li.lu.GetFixedText(toCode, deathMessageKey, translatedArgs)
	if err != nil {
		return "", err
	}
	return translatedMessage, nil
}
