package process

type Strategy interface {
	Process(arg interface{})
}
