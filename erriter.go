package erriter

import "iter"

type Seq[V any] func(errp *error) iter.Seq[V]
type ReturnErr[V any] func(yield func(V) bool) error

func New[V any](s ReturnErr[V]) Seq[V] {
	return func(errp *error) iter.Seq[V] {
		return func(yield func(V) bool) {
			if err := s(yield); err != nil {
				*errp = err
			}
		}
	}
}

type ReturnErr2[K, V any] func(yield func(K, V) bool) error
type Seq2[K, V any] func(errp *error) iter.Seq2[K, V]

func New2[K, V any](s ReturnErr2[K, V]) Seq2[K, V] {
	return func(errp *error) iter.Seq2[K, V] {
		return func(yield func(K, V) bool) {
			if err := s(yield); err != nil {
				*errp = err
			}
		}
	}
}
