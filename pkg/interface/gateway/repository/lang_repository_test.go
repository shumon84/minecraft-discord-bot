package repository

import (
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/shumon84/minecraft-discord-bot/pkg/core"
	"github.com/shumon84/minecraft-discord-bot/pkg/domain/entity"
)

func TestLangRepository_Find(t *testing.T) {
	type fields struct {
		resourcePack fs.FS
	}
	type args struct {
		langCode core.LangCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Lang
		wantErr bool
	}{
		{
			name: "正常系のテスト",
			fields: fields{
				resourcePack: fstest.MapFS{
					"minecraft/lang/ja_jp.json": &fstest.MapFile{Data: []byte(`
{
  "hoge" : "fuga",
  "foo" : "bar"
}
`)},
				},
			},
			args: args{
				langCode: core.JaJp,
			},
			want: entity.NewLang(core.JaJp, map[string]string{
				"hoge": "fuga",
				"foo":  "bar",
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LangRepository{
				resourcePack: tt.fields.resourcePack,
			}
			got, err := l.Find(tt.args.langCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}
