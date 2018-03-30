rerr
====

Golang helper for Gorilla JSON RPC 2.0. Common errors with codes and messages.

Not a big deal, just a little helper.

```go
if err != nil {
    rerr.New(rerr.EDuplicate, err)
}
```

Gorilla JSON RPC 2.0: <https://github.com/gorilla/rpc/tree/master/v2>

License
=======

MIT
