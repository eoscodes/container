package templates

type Container interface {
	Empty() bool
	Size() int
	Clear()

}

type Serializer interface {
	ToJSON() ([]byte, error)
	FromJSON([]byte) error
}

type Set interface {
	Container
	Serializer
}

type Map interface {
	Container
	Serializer
}

type A = int
