# rpchan

`rpchan` implements some of Go's channels semantics over a TCP connection using `net/rpc`.

It achieves this by providing a very minimal API on the `rpchan` type, exposing only three methods: `Send`, `Receive` and `Iter`.

It is advisible, but not mandatory, to use the same type on both the receiver and sender. This is because the types follow the [`encoding/gob` guidelines for types and values](https://pkg.go.dev/encoding/gob#hdr-Types_and_Values).