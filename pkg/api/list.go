package api

import (
	"io/fs"
	"time"

	"github.com/OlshaMB/modernandsimplefm/cmd/config"
	"github.com/spf13/afero"
)
var AppFs = afero.NewOsFs();

type Modified struct {
	By int
	At time.Time
}
type Entry struct {
	Name string
	Path string /// Should always be absolute
    CreatedAt time.Time
	Modified Modified
	IsDir bool
}
func ListDir(path string) ([]Entry, error){
	AppFolder, FolderErr := AppFs.Open(config.Cfg.Dir + path);
	if FolderErr != nil {
		return nil, FolderErr
	}
	entries := make([]Entry, 0)
	fsEntries, err := AppFolder.Readdir(0)
	if err != nil {
		return nil, err
	}
	for _, fsEntry := range fsEntries {
		entries = append(entries, Entry {
			Name: fsEntry.Name(),
			Path: path + "/" + fsEntry.Name(),
			CreatedAt: fsEntry.ModTime(),
			IsDir: fsEntry.IsDir(),
			Modified: Modified {
				By: 0,
				At: fsEntry.ModTime(),
			},
		})
	}
	AppFolder.Close()
	return entries, nil
}
func IsFile(path string) (bool, error){
	entry, err := AppFs.Stat(config.Cfg.Dir + path)
	if err!=nil {
		return false, err
	}
	return !entry.IsDir(), nil
}
func ReadFile(path string) (afero.File, error) {

	return AppFs.OpenFile(config.Cfg.Dir + path, 0, fs.FileMode(444))
}