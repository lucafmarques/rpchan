package receiver

import "errors"

// Receiver needs to be exported to be used as an rpc handler.
type Receiver[T any] struct {
	channel chan *T
}

func NewReceiver[T any](buf uint) *Receiver[T] { return &Receiver[T]{make(chan *T, buf)} }

// Push implements the function signature for an rpc handler.
func (r *Receiver[T]) Send(item *T, ok *bool) error {
	var err error
	defer func() {
		recover()
		// safe to assume that we're recovering from a send on a closed channel
		err = errors.New("queue is closed")
	}()

	r.channel <- item

	return err
}

func (r Receiver[T]) Close()           { close(r.channel) }
func (r Receiver[T]) Channel() chan *T { return r.channel }
