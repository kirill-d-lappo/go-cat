package iterator


type arrayCollection struct {
	source []interface{}
}

type arrayIterator struct {
	source       []interface{}
	currentIndex int
}

func (i *arrayIterator) Move() (bool, error) {
	i.currentIndex += 1
	if i.currentIndex >= len(i.source){
		return false, nil
	}

	return true, nil
}

func (i *arrayIterator) Current() interface{} {
	index := i.currentIndex
	source := i.source
	if index < 0 || index >= len(source) {
		return nil
	}

	return source[index]
}

func (r *arrayCollection) GetIterator() Iterator {
	return &arrayIterator{
		source:       r.source,
		currentIndex: -1,
	}
}

func NewIterable(source []interface{}) Iterable {
	return &arrayCollection{ source: source}
}
