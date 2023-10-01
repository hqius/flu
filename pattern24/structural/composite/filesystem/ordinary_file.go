package filesystem

type OrdFile interface {
	File
	Content() []byte
}
