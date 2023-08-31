# rpchan

`rpchan` implements some of Go's channels semantics over a TCP connection using `net/rpc`.

It achieves this by providing a very minimal API, consisting of only two functions, `Send` and `Receive`.

Because we can't really mimic the type safety of send and receive channel operations on a TCP connection, the type used for instantiating `Receive` must be the same type of the data sent using `Send`.