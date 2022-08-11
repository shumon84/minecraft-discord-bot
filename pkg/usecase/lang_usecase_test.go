package usecase

import (
	"reflect"
	"testing"

	"github.com/shumon84/minecraft-discord-bot/pkg/core"
	"github.com/shumon84/minecraft-discord-bot/pkg/domain/entity"
)

func TestLangUsecase_DeathMessageToLangKeyAndArgs(t *testing.T) {
	type fields struct {
		lr LangRepository
	}
	type args struct {
		message  string
		langCode core.LangCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   []string
		wantErr bool
	}{
		{
			name: "正常系のテスト",
			fields: fields{
				lr: &MockLangRepository{
					FindExpectBehavior: func(langCode core.LangCode) (*entity.Lang, error) {
						if langCode != core.EnGb {
							panic("invalid")
						}
						return entity.NewLang(core.EnGb, map[string]string{
							"death.attack.arrow": "%1$s was shot by %2$s",
						}), nil
					},
				},
			},
			args: args{
				message:  "shumon_84 was shot by Skeleton Horse",
				langCode: core.EnGb,
			},
			want:    "death.attack.arrow",
			want1:   []string{"shumon_84", "Skeleton Horse"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lu := &LangUsecase{
				lr: tt.fields.lr,
			}
			got, got1, err := lu.DeathMessageToLangKeyAndArgs(tt.args.message, tt.args.langCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeathMessageToLangKeyAndArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeathMessageToLangKeyAndArgs() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DeathMessageToLangKeyAndArgs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLangUsecase_EntityToLangKeyAndArgs(t *testing.T) {
	type fields struct {
		lr LangRepository
	}
	type args struct {
		message  string
		langCode core.LangCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   []string
		wantErr bool
	}{
		{
			name: "正常系のテスト",
			fields: fields{
				lr: &MockLangRepository{
					FindExpectBehavior: func(langCode core.LangCode) (*entity.Lang, error) {
						if langCode != core.EnGb {
							panic("invalid")
						}
						return entity.NewLang(core.EnGb, map[string]string{
							"entity.minecraft.skeleton_horse": "Skeleton Horse",
						}), nil
					},
				},
			},
			args: args{
				message:  "Skeleton Horse",
				langCode: core.EnGb,
			},
			want:    "entity.minecraft.skeleton_horse",
			want1:   []string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lu := &LangUsecase{
				lr: tt.fields.lr,
			}
			got, got1, err := lu.EntityToLangKeyAndArgs(tt.args.message, tt.args.langCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("EntityToLangKeyAndArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EntityToLangKeyAndArgs() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("EntityToLangKeyAndArgs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLangUsecase_GetFixedText(t *testing.T) {
	type fields struct {
		lr LangRepository
	}
	type args struct {
		langCode core.LangCode
		key      string
		args     []string
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
				lr: &MockLangRepository{
					FindExpectBehavior: func(langCode core.LangCode) (*entity.Lang, error) {
						if langCode != core.EnGb {
							panic("invalid")
						}
						return entity.NewLang(core.EnGb, map[string]string{
							"death.attack.arrow":              "%1$s was shot by %2$s",
							"entity.minecraft.skeleton_horse": "Skeleton Horse",
						}), nil
					},
				},
			},
			args: args{
				langCode: core.EnGb,
				key:      "death.attack.arrow",
				args:     []string{"shumon_84", "Skeleton Horse"},
			},
			want:    "shumon_84 was shot by Skeleton Horse",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lu := &LangUsecase{
				lr: tt.fields.lr,
			}
			got, err := lu.GetFixedText(tt.args.langCode, tt.args.key, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFixedText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFixedText() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestLangUsecase_textToLangKeyAndArgs(t *testing.T) {
//	type fields struct {
//		lr LangRepository
//	}
//	type args struct {
//		text     string
//		pattern  *regexp.Regexp
//		langCode core.LangCode
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    string
//		want1   []string
//		wantErr bool
//	}{
//		{
//			name: "正常系のテスト",
//			fields: fields{
//				lr: &MockLangRepository{
//					FindExpectBehavior: func(langCode core.LangCode) (*entity.Lang, error) {
//						switch langCode {
//						case core.JaJp:
//							return entity.NewLang(core.JaJp, map[string]string{
//								"death.attack.arrow":              "%1$sは%2$sに射抜かれた",
//								"entity.minecraft.skeleton_horse": "スケルトンホース",
//							}), nil
//						case core.EnGb:
//							return entity.NewLang(core.JaJp, map[string]string{
//								"death.attack.arrow":              "%1$s was shot by %2$s",
//								"entity.minecraft.skeleton_horse": "Skeleton Horse",
//							}), nil
//						default:
//							panic("invalid")
//						}
//					},
//				},
//			},
//			args:    args{},
//			want:    "",
//			want1:   nil,
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			lu := &LangUsecase{
//				lr: tt.fields.lr,
//			}
//			got, got1, err := lu.textToLangKeyAndArgs(tt.args.text, tt.args.pattern, tt.args.langCode)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("textToLangKeyAndArgs() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("textToLangKeyAndArgs() got = %v, want %v", got, tt.want)
//			}
//			if !reflect.DeepEqual(got1, tt.want1) {
//				t.Errorf("textToLangKeyAndArgs() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
