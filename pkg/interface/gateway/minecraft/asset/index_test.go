package asset

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestNewIndex(t *testing.T) {
	type args struct {
		indexFile io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *Index
		wantErr bool
	}{
		{
			name: "正常系のテスト",
			args: args{
				indexFile: bytes.NewReader([]byte(`
{
  "objects": {
    "icons/icon_16x16.png": {
      "hash": "bdf48ef6b5d0d23bbb02e17d04865216179f510a",
      "size": 3665
    },
    "icons/icon_32x32.png": {
      "hash": "92750c5f93c312ba9ab413d546f32190c56d6f1f",
      "size": 5362
    },
    "icons/minecraft.icns": {
      "hash": "991b421dfd401f115241601b2b373140a8d78572",
      "size": 114786
    }
  }
}
`)),
			},
			want: &Index{
				EntryDict: map[string]*IndexEntry{
					"icons/icon_16x16.png": {
						Hash: "bdf48ef6b5d0d23bbb02e17d04865216179f510a",
						Size: 3665,
					},
					"icons/icon_32x32.png": {
						Hash: "92750c5f93c312ba9ab413d546f32190c56d6f1f",
						Size: 5362,
					},
					"icons/minecraft.icns": {
						Hash: "991b421dfd401f115241601b2b373140a8d78572",
						Size: 114786,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIndex(tt.args.indexFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexEntry_Path(t *testing.T) {
	type fields struct {
		Hash string
		Size int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "正常系のテスト",
			fields: fields{
				Hash: "0123456789",
				Size: 0,
			},
			want: pathObjects + "/01/0123456789",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndexEntry{
				Hash: tt.fields.Hash,
				Size: tt.fields.Size,
			}
			if got := i.Path(); got != tt.want {
				t.Errorf("Path() = %v, want %v", got, tt.want)
			}
		})
	}
}
