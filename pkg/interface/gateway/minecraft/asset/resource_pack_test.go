package asset

import (
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func HelperGenerateTestMinecraftFS(t *testing.T) fs.FS {
	t.Helper()
	return fstest.MapFS{
		"assets/indexes/1.19.json": &fstest.MapFile{Data: []byte(`
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
`)},
		"assets/objects/bd/bdf48ef6b5d0d23bbb02e17d04865216179f510a": &fstest.MapFile{Data: make([]byte, 3665)},
		"assets/objects/92/92750c5f93c312ba9ab413d546f32190c56d6f1f": &fstest.MapFile{Data: make([]byte, 5362)},
		"assets/objects/99/991b421dfd401f115241601b2b373140a8d78572": &fstest.MapFile{Data: make([]byte, 114786)},
	}

}

func TestNewResourcePack(t *testing.T) {
	testFS := HelperGenerateTestMinecraftFS(t)

	type args struct {
		minecraftPath fs.FS
		version       string
	}
	tests := []struct {
		name    string
		args    args
		want    *ResourcePack
		wantErr bool
	}{
		{
			name: "正常系のテスト",
			args: args{
				minecraftPath: testFS,
				version:       "1.19",
			},
			want: &ResourcePack{
				version:     "1.19",
				minecraftFS: testFS,
				index: &Index{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewResourcePack(tt.args.minecraftPath, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewResourcePack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResourcePack() got = %v, want %v", got, tt.want)
			}
		})
	}
}
