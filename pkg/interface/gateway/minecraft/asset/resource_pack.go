package asset

import (
	"io/fs"
	"path/filepath"
)

const (
	pathAssets  = "assets"
	pathIndexes = pathAssets + string(filepath.Separator) + "indexes"
	pathObjects = pathAssets + string(filepath.Separator) + "objects"
)

type ResourcePack struct {
	version     string
	minecraftFS fs.FS // .minecraft/
	index       *Index
}

var _ = fs.FS(new(ResourcePack))

func NewResourcePack(minecraftPath fs.FS, version string) (*ResourcePack, error) {
	indexFilePath := filepath.Join(pathIndexes, version+".json")
	indexFile, err := minecraftPath.Open(indexFilePath)
	if err != nil {
		return nil, err
	}
	defer indexFile.Close()
	index, err := NewIndex(indexFile)
	if err != nil {
		return nil, err
	}
	return &ResourcePack{
		version:     version,
		minecraftFS: minecraftPath,
		index:       index,
	}, nil
}

func (r *ResourcePack) Open(name string) (fs.File, error) {
	entry := r.index.EntryDict[name]
	return r.minecraftFS.Open(entry.Path())
}
