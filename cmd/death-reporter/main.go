package main

import (
	"log"
	"os"

	"github.com/shumon84/minecraft-discord-bot/pkg/infrastructure"
	"github.com/shumon84/minecraft-discord-bot/pkg/infrastructure/config"
	"github.com/shumon84/minecraft-discord-bot/pkg/interface/controller"
	"github.com/shumon84/minecraft-discord-bot/pkg/interface/gateway/discord"
	"github.com/shumon84/minecraft-discord-bot/pkg/interface/gateway/minecraft/asset"
	"github.com/shumon84/minecraft-discord-bot/pkg/interface/gateway/repository"
	"github.com/shumon84/minecraft-discord-bot/pkg/interface/interactor"
	"github.com/shumon84/minecraft-discord-bot/pkg/usecase"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	rp, err := asset.NewResourcePack(os.DirFS(cfg.Minecraft.Path), cfg.Minecraft.Version)
	if err != nil {
		log.Fatal(err)
	}
	lp := repository.NewLangRepository(rp)
	lu := usecase.NewLangUsecase(lp)
	li := interactor.NewLangInteractor(lu)
	lc := controller.NewLangController(li)
	ls := infrastructure.NewLogScanner(os.Stdin)
	token := cfg.Discord.Token
	disco, err := discord.NewClient(token)
	if err != nil {
		log.Fatal(err)
	}
	channelID := cfg.Discord.ChannelID
	logger := log.Default()
	daemon := infrastructure.NewDeathReporterDaemon(ls, lc, disco, channelID, logger)
	daemon.Run()
}
