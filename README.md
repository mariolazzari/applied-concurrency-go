# Applied concurrency in Go

## Go routines

- Independently executing functions
- Lightweight thread
- On top of threads

### Parallelism

Parallel events or tasks execute simultaneously and independently. 
True parallel events require multimple CPUs.

### Concurrency

Concurrent tasks or events are interleaving and can happen in any given order. 
It is a non deterministic way of achieving multiple tasks.

### Wait group

- Wait until go routines finish
- Counter and inner lock
- Zero value ready
- sync.WaitGroup
  - Add
  - Done
  - Wait
- All the types in sync package must be passed as pointers

[Sync](https://pkg.go.dev/sync)

### Race conditions

Use -race flag in order to detect race conditions.

### sync.Map

- safe for go routines
- performance overhead

### sync.Mutex

- initialized unlocled
- Lock
- Unlock