[![Go Reference](https://pkg.go.dev/badge/github.com/lucafmarques/rpchan.svg)](https://pkg.go.dev/github.com/lucafmarques/rpchan)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucafmarques/rpchan)](https://goreportcard.com/report/github.com/lucafmarques/rpchan)

# rpchan

`rpchan` implements some of Go's channels semantics over a TCP connection using `net/rpc`.

It achieves this by providing a very minimal API on the `RPChan` type, exposing 4 methods: `Send`, `Receive`, `Listen` and `Close`. Those four methods are enough to mimic one-way send/receive channel semantics.

It is advisable, but not mandatory, to use the same type on both the receiver and sender. This is because the types follow the [`encoding/gob` guidelines for types and values](https://pkg.go.dev/encoding/gob#hdr-Types_and_Values).

If built with `go1.22` and `GOEXPERIMENT=rangefunc`, the `Listen` method can be used on a for-range loop, working exactly like a Go channel would.
