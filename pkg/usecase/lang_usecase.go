package usecase

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/shumon84/minecraft-discord-bot/pkg/core"
	"github.com/shumon84/minecraft-discord-bot/pkg/domain/entity"
	"github.com/shumon84/minecraft-discord-bot/pkg/interface/interactor"
)

type LangUsecase struct {
	lr LangRepository
}

func NewLangUsecase(lr LangRepository) *LangUsecase {
	return &LangUsecase{
		lr: lr,
	}
}

func (lu *LangUsecase) textToLangKeyAndArgs(text string, pattern *regexp.Regexp, langCode core.LangCode) (string, []string, error) {
	lang, err := lu.lr.Find(langCode)
	if err != nil {
		return "", nil, err
	}

	deathMessages := lang.GetAll(pattern)
	for key, deathMessage := range deathMessages {
		args, err := deathMessage.GetArgs(text)
		target := &entity.ArgExtractNotMatchError{}
		if errors.As(err, &target) {
			// マッチしなかった場合、別のプレースホルダとマッチするか試すために continue
			continue
		}
		if err != nil {
			return "", nil, err
		}
		return key, args, nil
	}
	return "", nil, fmt.Errorf("%w: %s", interactor.ErrLangUseCaseNotFoundMatchedPlaceHolder, text)
}

func (lu *LangUsecase) DeathMessageToLangKeyAndArgs(message string, langCode core.LangCode) (string, []string, error) {
	return lu.textToLangKeyAndArgs(message, regexp.MustCompile(`^death\.`), langCode)
}

func (lu *LangUsecase) EntityToLangKeyAndArgs(message string, langCode core.LangCode) (string, []string, error) {
	return lu.textToLangKeyAndArgs(message, regexp.MustCompile(`^entity\.`), langCode)
}

func (lu *LangUsecase) GetFixedText(langCode core.LangCode, key string, args []string) (string, error) {
	lang, err := lu.lr.Find(langCode)
	if err != nil {
		return "", err
	}
	placeHolder, err := lang.Get(key)
	if err != nil {
		return "", err
	}
	fixedText, err := placeHolder.Apply(args)
	if err != nil {
		return "", err
	}
	return fixedText, nil
}
