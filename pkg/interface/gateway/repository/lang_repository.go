package repository

import (
	"encoding/json"
	"io/fs"

	"github.com/shumon84/minecraft-discord-bot/pkg/core"
	"github.com/shumon84/minecraft-discord-bot/pkg/domain/entity"
	"github.com/shumon84/minecraft-discord-bot/pkg/usecase"
)

type LangRepository struct {
	resourcePack fs.FS
}

var _ = usecase.LangRepository(new(LangRepository))

const langPath = "minecraft/lang/"

func NewLangRepository(resourcePack fs.FS) *LangRepository {
	return &LangRepository{
		resourcePack: resourcePack,
	}
}

func (l *LangRepository) Find(langCode core.LangCode) (*entity.Lang, error) {
	langFile, err := l.resourcePack.Open(langPath + string(langCode) + ".json")
	if err != nil {
		return nil, err
	}
	defer langFile.Close()
	decoder := json.NewDecoder(langFile)
	dict := map[string]string{}
	if err := decoder.Decode(&dict); err != nil {
		return nil, err
	}
	return entity.NewLang(langCode, dict), nil
}
