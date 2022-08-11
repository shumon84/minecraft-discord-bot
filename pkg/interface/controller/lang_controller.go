package controller

import (
	"github.com/shumon84/minecraft-discord-bot/pkg/core"
	"github.com/shumon84/minecraft-discord-bot/pkg/interface/gateway/minecraft/console"
	"github.com/shumon84/minecraft-discord-bot/pkg/interface/interactor"
)

type LangController struct {
	li *interactor.LangInteractor
}

func NewLangController(li *interactor.LangInteractor) *LangController {
	return &LangController{
		li: li,
	}
}

func (lc *LangController) TranslateMessage(input *console.Log) (string, error) {
	translated, err := lc.li.TranslateDeathMessage(input.Message, core.EnGb, core.JaJp)
	if err != nil {
		return "", err
	}
	return translated, nil
}
