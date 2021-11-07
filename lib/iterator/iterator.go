package iterator

type Iterator interface {
	Move() (bool, error)
	Current() interface{}
}

type Iterable interface {
	GetIterator() Iterator
}