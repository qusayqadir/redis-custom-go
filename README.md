# redis-custom-go

Building my own Redis clone from scratch in Go.

## to do 

### `store.go` — the data layer

- [x] `Set`, `Get`, `Exist`, `Del` with unit tests in CI
- [ ] Add a `sync.RWMutex` to `Store` for safe concurrent access
- [ ] Add TTL: `Expire`, `TTL`, and `Persist`
- [ ] Expire keys lazily on `Get`/`Exist`

### `resp.go` — the wire protocol

- [ ] `Decode` a RESP array of bulk strings into a command
- [ ] Encode replies: simple string, error, integer, bulk string, null
- [ ] Unit tests for decoding and each encoder

### `server.go` — networking + dispatch

- [x] TCP listener that spawns a goroutine per client
- [ ] Implement `handleConn` to decode, dispatch, and reply in a loop
- [ ] Dispatch `PING`, `SET`, `GET`, `DEL`, `EXISTS`, `EXPIRE`, `TTL`
