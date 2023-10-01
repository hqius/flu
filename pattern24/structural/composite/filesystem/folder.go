package filesystem

type Folder interface {
	File
	Children() []File
}
