package iterator

type Iterator[T any] interface {
	IsDone() bool
	CurrentItem() T
	Next()
	Reset()
}
