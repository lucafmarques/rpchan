# rpchan

`rpchan` implements some of Go's channels semantics over a TCP connection using `net/rpc`.

It achieves this by providing a very minimal API on the `rpchan` type, exposing only three methods: `Send`, `Receive` and `Iter`.

Because we can't really mimic the type safety of send and receive channel operations on a TCP connection, the type of a receive rpchan must be as similar as possible to as that of a send rpchan.
This means that the types used have follow the [`encoding/gob` guidelines for types and values](https://pkg.go.dev/encoding/gob#hdr-Types_and_Values).