package producer

import (
	"net/rpc"
)

type producerer interface {
	Send(item any) (ok bool)
}

type producer struct {
	client *rpc.Client
}

func (p *producer) Send(t any) (ok bool) {
	p.client.Call("RPCReceiver.Receive", t, ok)
	return
}

var (
	c, err   = rpc.Dial("tcp", ":1001")
	Producer = producer{c}
)
