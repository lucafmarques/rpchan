package receiver

var t = true

type Receiver[T any] struct {
	Queue chan *T
}

func (r *Receiver[T]) Receive(item *T, ok *bool) error {
	r.Queue <- item
	ok = &t
	return nil
}
