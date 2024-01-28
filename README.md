# `rpchan`: channel-like semantics over net/rpc
[![Go Reference](https://pkg.go.dev/badge/github.com/lucafmarques/rpchan.svg)](https://pkg.go.dev/github.com/lucafmarques/rpchan)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucafmarques/rpchan)](https://goreportcard.com/report/github.com/lucafmarques/rpchan)

`rpchan` implements some of Go's channel semantics over a TCP connection using `net/rpc`.

```
go get github.com/lucafmarques/rpchan
```

It achieves this by providing a minimal API on the `RPChan[T any]` type: 
- Use [`Send`](https://pkg.go.dev/github.com/lucafmarques/rpchan#RPChan.Send) if you want to send a `T`, similarly to `ch <- T`
- Use [`Receive`](https://pkg.go.dev/github.com/lucafmarques/rpchan#RPChan.Receive) if you want to receive a `T`, like `<-ch`
- Use [`Close`](https://pkg.go.dev/github.com/lucafmarques/rpchan#RPChan.Close) if you want to close the channel, like `close(ch)`
- Use [`Listen`](#rangefunc) if want to iterate on the channel, like `for v := range ch`

Those four methods are enough to mimic one-way send/receive channel-like semantics.

It's advisable, but not mandatory, to use the same type on both the receiver and sender. This is because `rpchan` follows the [`encoding/gob`](https://pkg.go.dev/encoding/gob#hdr-Types_and_Values) guidelines for encoding types and values.

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

---

README.md heavily inspired by [`sourcegraph/conc`](https://github.com/sourcegraph/conc)
