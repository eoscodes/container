package templates


type Container interface {
	Empty() bool
	Size() int
	Clear()

}

///go:generate go test -v github.com/ncw/gotemplate/...
//go:generate gotemplate "github.com/eosspark/container/templates/treeset" AccountNameSet(common.AccountName,common.CompareName)


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
