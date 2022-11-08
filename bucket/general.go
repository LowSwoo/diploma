package bucket

func newFolder(name string) *Folder {
	return &Folder{Name: name, Folders: make(map[string]Folder), Files: make(map[string]File)}
}

func NewBucket(name string) *Bucket {
	return &Bucket{Name: name, Folders: make(map[string]Folder)}
}
