package receiver

import "errors"

var t = true

type Receiver[T any] struct {
	Queue chan *T
}

func (r *Receiver[T]) Receive(item *T, ok *bool) error {
	var err error
	defer func() {
		recover()
		err = errors.New("queue is closed")
	}()

	r.Queue <- item
	ok = &t

	return err
}
