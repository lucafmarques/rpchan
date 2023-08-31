package receiver

import "errors"

var t = true

type Receiver[T any] struct {
	Queue chan *T
	err   error
}

func (r *Receiver[T]) Receive(item *T, ok *bool) error {
	defer func() {
		recover()
		r.err = errors.New("queue is closed")
	}()

	r.Queue <- item
	ok = &t

	return r.err
}
