package bucket

import (
	"strings"
)

type Folder struct {
	Name    string            `json:"name"`
	Folders map[string]Folder `json:"folders"`
	Files   map[string]File   `json:"files"`
}

func (f *Folder) New(name string) {
	f.Folders[name] = Folder{Name: name, Folders: make(map[string]Folder), Files: make(map[string]File)}
}

func (f *Folder) AppendFile(name, url string, info interface{}) {
	fs := strings.Split(name, "/")
	if len(fs) > 1 {
		if !f.checkFolder(fs[0]) {
			f.New(fs[0])
		}
		temp := f.Folders[fs[0]]
		temp.AppendFile(strings.Join(fs[1:], "/"), url, info)

	} else {
		f.Files[fs[0]] = File{Name: name, Link: url, Info: info}
	}
}

func (f *Folder) checkFolder(name string) bool {
	_, ex := f.Folders[name]
	return ex
}
