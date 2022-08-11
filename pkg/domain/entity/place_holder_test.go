package entity

import (
	"reflect"
	"testing"
)

func TestNewPlaceHolder(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want *PlaceHolder
	}{
		{
			name: "正常系のテスト",
			args: args{
				src: "%1$s was shot by %2$s using %3$s",
			},
			want: &PlaceHolder{
				format:    "%s was shot by %s using %s",
				numOfArgs: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlaceHolder(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlaceHolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaceHolder_Apply(t *testing.T) {
	type fields struct {
		format    string
		numOfArgs int
	}
	type args struct {
		args []string
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
				format:    "%s was shot by %s using %s",
				numOfArgs: 3,
			},
			args: args{
				args: []string{"PlayerA", "PlayerB", "WeaponA"},
			},
			want:    "PlayerA was shot by PlayerB using WeaponA",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlaceHolder{
				format:    tt.fields.format,
				numOfArgs: tt.fields.numOfArgs,
			}
			got, err := p.Apply(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Apply() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_argExtractor_ExtractArgs(t *testing.T) {
	type fields struct {
		formatWords      []string
		formedWords      []string
		formatWordsIndex int
		formedWordsIndex int
		result           []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "正常系の場合",
			fields: fields{
				formatWords:      []string{"%s", "was", "shot", "by", "%s"},
				formedWords:      []string{"PlayerA", "was", "shot", "by", "Skeleton"},
				formatWordsIndex: 0,
				formedWordsIndex: 0,
				result:           make([]string, 0, 2),
			},
			want:    []string{"PlayerA", "Skeleton"},
			wantErr: false,
		},
		{
			name: "正常系の場合2",
			fields: fields{
				formatWords:      []string{"%s", "was", "shot", "by", "%s"},
				formedWords:      []string{"PlayerA", "was", "shot", "by", "Wither", "Skeleton"},
				formatWordsIndex: 0,
				formedWordsIndex: 0,
				result:           make([]string, 0, 2),
			},
			want:    []string{"PlayerA", "Wither Skeleton"},
			wantErr: false,
		},
		{
			name: "マッチしない場合",
			fields: fields{
				formatWords:      []string{"%s", "was", "shot", "by", "%s"},
				formedWords:      []string{"Skeleton"},
				formatWordsIndex: 0,
				formedWordsIndex: 0,
				result:           make([]string, 0, 2),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &argExtractor{
				formatWords:      tt.fields.formatWords,
				formedWords:      tt.fields.formedWords,
				formatWordsIndex: tt.fields.formatWordsIndex,
				formedWordsIndex: tt.fields.formedWordsIndex,
				result:           tt.fields.result,
			}
			got, err := a.ExtractArgs()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
