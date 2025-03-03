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

## Channels

- No need to share memory
- Channel is a pass-through
- chan reserved keyword
- <-
- Associated to a type
- Zero value nil
- c := make(chan T)
- Sending data with a channel: ch <- data
- Receving data with a channel: data := <- ch
- Both operations are blocking
- Unbuffered channel: by default does not store value and need both sender and receaver
- Buffered channel: stores value for later use

### Channels direction

- Bidirectional: chan T
- Send only: chan <- T
- Receive only: <- chan T
- Bidirectional implicity cast to uni

### Close channel

- No more values will be sent after closing 
- close(ch)

### Select

- wait for multiple channel operations
- blocks until one operation is ready
- similar to switch