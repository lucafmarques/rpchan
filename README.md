# rpc-queue

`rpc-queue` aims to be a quick and dirty way of sending gob-encoded data to a sidecar worker.

It achieves this by providing a very minimal API, consisting of only two functions: `Send` and `Listen`.

There are only two rules you must follow:
- Data sent through `Send` must be of the same type a caller defines their `Listen` on.
- A sender must not listen and a listener must not send.