package filesystem

type File interface {
	Name() string
	Type() FileType
}
