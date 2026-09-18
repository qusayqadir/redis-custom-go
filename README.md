# redis-custom-go

Building my own Redis clone from scratch in Go.

## What this is

The goal is to reimplement Redis myself while learning Go. 
start with a single-node server that runs on my machine, then to grow it into something
distributed. 

## Roadmap

### Phase 1 — Local Redis clone

Get a real, working key-value server running locally:

- [ ] TCP server that accepts client connections (`net` package)
- [ ] Speak the RESP protocol (the wire format `redis-cli` uses)
- [ ] Core commands: `PING`, `SET`, `GET`, `DEL`
- [ ] Key expiry (`EX` / `TTL`)
- [ ] Concurrency: handle many clients at once with goroutines

Success = I can point `redis-cli` at it and it just works.

### Phase 2 — Redis "cloud" (distributed)

Once the single node is solid, make it distributed and design the cloud
architecture myself:

- [ ] Replication (primary → replicas)
- [ ] Sharding / partitioning keys across nodes
- [ ] A coordination/routing layer
- [ ] Cloud architecture design (deployment, scaling, failover)



## Why

The point isn't to replace Redis — it's to **learn Go**, understand how Redis
works by building it, and see how far I can take it.

## Getting started

Requires [Go](https://go.dev/dl/) installed.

```bash
# Run it
go run .

# Or build a binary
go build -o redis-custom-go .
./redis-custom-go
```

