package utils

type Comparable interface {
	Compare(c Comparable) int
}
