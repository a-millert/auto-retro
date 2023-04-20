package misc

type Set[T comparable] map[T]bool

func Deduplicate[T comparable](xs []T) Set[T] {
	set := make(Set[T])
	for _, x := range xs {
		set[x] = true
	}
	return set
}
