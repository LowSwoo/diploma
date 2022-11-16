package bucket

import (
	"errors"
	"fmt"
	"strings"
)

type Bucket struct {
	Name        string            `json:"bucket"`
	FoldersList []string          `json:"foldersList"`
	Folders     map[string]Folder `json:"folders"`
}

func (b *Bucket) AppendFile(name, url string, info interface{}) error {
	fs := strings.Split(name, "/")
	if len(fs) <= 1 {
		return errors.New("В бакете не может быть файлов")
	}
	if !b.isFolder(fs[0]) {
		b.Folders[fs[0]] = *newFolder(fs[0])
	}
	temp := b.Folders[fs[0]]
	temp.AppendFile(strings.Join(fs[1:], "/"), url, info)
	return nil
}

func (b *Bucket) SetFoldersList() {
	keys := make([]string, 0, len(b.Folders))
	for k := range b.Folders {
		keys = append(keys, k)
	}
	b.FoldersList = keys
	fmt.Printf("%v", b.FoldersList)

}

func (b *Bucket) isFolder(name string) bool {
	_, ex := b.Folders[name]
	return ex
}
