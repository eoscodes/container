package templates


type Container interface {
	Empty() bool
	Size() int
	Clear()

	Serializer
}

type Serializer interface {
	ToJSON() ([]byte, error)
	FromJSON([]byte) error
}

type Set interface {
	Container
}

type Map interface {
	Container
}
