package asset

import (
	"encoding/json"
	"io"
	"path/filepath"
)

type Index struct {
	EntryDict map[string]*IndexEntry `json:"objects"`
}

type IndexEntry struct {
	Hash string `json:"hash"`
	Size int    `json:"size"`
}

func (i *IndexEntry) Path() string {
	return filepath.Join(pathObjects, string([]rune(i.Hash)[:2]), i.Hash)
}

func NewIndex(indexFile io.Reader) (*Index, error) {
	index := &Index{}
	decoder := json.NewDecoder(indexFile)
	if err := decoder.Decode(index); err != nil {
		return nil, err
	}
	return index, nil
}
