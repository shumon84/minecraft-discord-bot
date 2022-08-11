package entity

import (
	"fmt"
	"regexp"

	"github.com/shumon84/minecraft-discord-bot/pkg/core"
)

type LangNotFoundError struct {
	lang *Lang
	key  string
}

func (l *LangNotFoundError) Error() string {
	return fmt.Sprintf("a text correspond to \"%s\" is not found in the %s dictionary", l.key, l.lang.LangCode)
}

func (l *Lang) notFoundError(key string) error {
	return &LangNotFoundError{
		lang: l,
		key:  key,
	}
}

type LangNotFoundKeyError struct {
	lang *Lang
	text string
}

func (l *LangNotFoundKeyError) Error() string {
	return fmt.Sprintf("a key correspond to \"%s\" is not found in the %s dictionary", l.text, l.lang.LangCode)
}

func (l *Lang) notFoundKeyError(text string) error {
	return &LangNotFoundKeyError{
		lang: l,
		text: text,
	}
}

type Lang struct {
	LangCode core.LangCode
	dict     map[string]*PlaceHolder
}

func NewLang(langCode core.LangCode, textDict map[string]string) *Lang {
	placeHolderDict := map[string]*PlaceHolder{}
	for key, text := range textDict {
		placeHolderDict[key] = NewPlaceHolder(text)
	}
	return &Lang{
		LangCode: langCode,
		dict:     placeHolderDict,
	}
}

func (l *Lang) Get(key string) (*PlaceHolder, error) {
	placeHolder, ok := l.dict[key]
	if !ok {
		return nil, l.notFoundError(key)
	}
	return placeHolder, nil
}
func (l *Lang) MustGet(key string) *PlaceHolder {
	placeHolder, err := l.Get(key)
	if err != nil {
		panic(err)
	}
	return placeHolder
}

func (l *Lang) GetAll(keyPattern *regexp.Regexp) map[string]*PlaceHolder {
	matched := map[string]*PlaceHolder{}
	for key, placeHolder := range l.dict {
		if keyPattern.MatchString(key) {
			matched[key] = placeHolder
		}
	}
	return matched
}

func (l *Lang) FindKey(text string) (string, error) {
	for key, placeHolder := range l.dict {
		if placeHolder.format == text {
			return key, nil
		}
	}
	return "", l.notFoundKeyError(text)
}
