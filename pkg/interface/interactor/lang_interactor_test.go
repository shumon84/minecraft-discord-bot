package interactor

import (
	"fmt"
	"testing"

	"github.com/shumon84/minecraft-discord-bot/pkg/core"
)

func TestLangInteractor_TranslateDeathMessage(t *testing.T) {
	type fields struct {
		lu LangUsecase
	}
	type args struct {
		message  string
		fromCode core.LangCode
		toCode   core.LangCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "正常系のテスト",
			fields: fields{
				lu: &MockLangUsecase{
					DeathMessageToLangKeyAndArgsExpectBehavior: func(message string, langCode core.LangCode) (string, []string, error) {
						if message != "shumon_84 was shot by Skeleton Horse" || langCode != core.EnGb {
							panic("invalid")
						}
						return "death.attack.arrow", []string{"shumon_84", "Skeleton Horse"}, nil
					},
					EntityToLangKeyAndArgsExpectBehavior: func(message string, langCode core.LangCode) (string, []string, error) {
						if message != "Skeleton Horse" || langCode != core.EnGb {
							return "", nil, ErrLangUseCaseNotFoundMatchedPlaceHolder
						}
						return "entity.minecraft.skeleton_horse", []string{}, nil
					},
					GetFixedTextExpectBehavior: func(langCode core.LangCode, key string, args []string) (string, error) {
						if langCode != core.JaJp {
							panic("invalid")
						}
						switch key {
						case "death.attack.arrow":
							if len(args) != 2 {
								panic("invalid")
							}
							return fmt.Sprintf("%sは%sに射抜かれた", args[0], args[1]), nil
						case "entity.minecraft.skeleton_horse":
							if len(args) != 0 {
								panic("invalid")
							}
							return "スケルトンホース", nil
						default:
							panic("invalid")
						}
					},
				},
			},
			args: args{
				message:  "shumon_84 was shot by Skeleton Horse",
				fromCode: core.EnGb,
				toCode:   core.JaJp,
			},
			want:    "shumon_84はスケルトンホースに射抜かれた",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			li := &LangInteractor{
				lu: tt.fields.lu,
			}
			got, err := li.TranslateDeathMessage(tt.args.message, tt.args.fromCode, tt.args.toCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("TranslateDeathMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TranslateDeathMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
