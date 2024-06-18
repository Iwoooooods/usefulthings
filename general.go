package collection

type GeneralFunc[T comparable] interface {
	Contains(elem T) bool
	Add(elem T)
	Remove(elem T)
	Size() int
	Iterate() <-chan T
}
