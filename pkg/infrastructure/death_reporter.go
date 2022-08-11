package infrastructure

import (
	"log"

	"github.com/shumon84/minecraft-discord-bot/pkg/interface/controller"
	"github.com/shumon84/minecraft-discord-bot/pkg/interface/gateway/discord"
	"github.com/shumon84/minecraft-discord-bot/pkg/interface/gateway/minecraft/console"
)

type DeathReporterDaemon struct {
	ls        *LogScanner
	lc        *controller.LangController
	disco     discord.Client
	channelID string
	logger    *log.Logger
}

func NewDeathReporterDaemon(ls *LogScanner, lc *controller.LangController, disco discord.Client, channelID string, logger *log.Logger) *DeathReporterDaemon {
	return &DeathReporterDaemon{
		ls:        ls,
		lc:        lc,
		disco:     disco,
		channelID: channelID,
		logger:    logger,
	}
}

func (d *DeathReporterDaemon) Run() {
	for d.ls.Scan() {
		go d.run(d.ls.Log())
	}
}

func (d *DeathReporterDaemon) run(input *console.Log) {
	defer func() {
		err := recover()
		if err != nil {
			d.logger.Println("[FATAL] ", err)
		}
	}()
	if input.LogLevel != console.LogLevelInfo {
		return
	}
	if input.ThreadName != "Server thread" {
		return
	}
	translated, err := d.lc.TranslateMessage(d.ls.Log())
	if err != nil {
		d.logger.Println("[ERROR] ", err)
		return
	}
	if err := d.disco.ChannelMessageSend(d.channelID, translated); err != nil {
		d.logger.Println("[ERROR] ", err)
	}
}
