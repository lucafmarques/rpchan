# `rpchan`: channel-like semantics over net/rpc. 
[![Go Reference](https://pkg.go.dev/badge/github.com/lucafmarques/rpchan.svg)](https://pkg.go.dev/github.com/lucafmarques/rpchan)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucafmarques/rpchan)](https://goreportcard.com/report/github.com/lucafmarques/rpchan)

`rpchan` implements some of Go's channel semantics over a TCP connection using `net/rpc`.

It achieves this by providing a very minimal API on the `RPChan` type, exposing four methods: `Send`, `Receive`, [`Listen`](#rangefunc) and `Close`. Those four methods are enough to mimic one-way send/receive channel-like semantics.

It is advisable, but not mandatory, to use the same type on both the receiver and sender. This is because the types follow the [`encoding/gob` guidelines for types and values](https://pkg.go.dev/encoding/gob#hdr-Types_and_Values).

## Examples
<table>
<tr>
<th><code>sender.go</code></th>
<th><code>receiver.go</code></th>
</tr>
<tr>
<td>
  
```go
package main

import (
	"github.com/lucafmarques/rpchan"
)

func main() {
	ch := rpchan.New[int](":9091")
	err := ch.Send(20)
	// error handling because of
	// the required network call
	err = ch.Close()
}
```
</td>
<td>
  
```go
package main

import (
	"github.com/lucafmarques/rpchan"
)

func main() {
	ch := rpchan.New[int](":9091", 100)
	for v := range ch.Listen() {
		// ...
	}
}
```
</td>
</tr>
</table>

## rangefunc
If built with Go 1.22 and `GOEXPERIMENT=rangefunc`, the `Listen` method can be used on a for-range loop, working exactly like a Go channel would.
