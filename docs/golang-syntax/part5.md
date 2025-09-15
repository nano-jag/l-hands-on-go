## Part 5: Concurrency (goroutines, channels, sync, context)

### Goroutines

```go
go work() // lightweight thread; do not leak
```

Guidelines:
- Launch with clear ownership and lifetime; avoid goroutine leaks.
- Capture loop vars correctly (shadow inside loop).

### Channels

```go
ch := make(chan int)      // unbuffered
ch2 := make(chan int, 16) // buffered

go func() { defer close(ch); for i := 0; i < 3; i++ { ch <- i } }()
for v := range ch { _ = v }
```

Rules:
- Send on closed channel panics; receive from closed yields zero, ok=false in two-value receive.
- Close by the sender only, to signal "no more values".

### Select

```go
select {
case v := <-ch:
    _ = v
case ch <- 42:
    // sent
case <-time.After(100 * time.Millisecond):
    // timeout
}
```

### Context for cancellation and deadlines

```go
func fetch(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    case <-time.After(50 * time.Millisecond):
        return nil
    }
}

ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
_ = fetch(ctx)
```

Guidelines:
- First arg is `context.Context`; do not store contexts in structs.
- Always cancel to free resources.

### sync primitives

```go
var mu sync.Mutex
mu.Lock(); /* critical */; mu.Unlock()

var rw sync.RWMutex
rw.RLock(); /* read */; rw.RUnlock()
rw.Lock(); /* write */; rw.Unlock()

var wg sync.WaitGroup
wg.Add(n)
for i := 0; i < n; i++ { go func(){ defer wg.Done() }() }
wg.Wait()

var once sync.Once
once.Do(func(){ /* init */ })

var v atomic.Int64
v.Add(1)
```

Guidelines:
- Do not copy structs containing mutexes/cond once in use.
- Prefer `sync/atomic` for high-frequency counters; use mutex for compound invariants.

### Worker pool pattern

```go
type Task func() error

func RunPool(ctx context.Context, n int, tasks <-chan Task) error {
    g, ctx := errgroup.WithContext(ctx)
    for i := 0; i < n; i++ {
        g.Go(func() error {
            for {
                select {
                case <-ctx.Done():
                    return ctx.Err()
                case t, ok := <-tasks:
                    if !ok { return nil }
                    if err := t(); err != nil { return err }
                }
            }
        })
    }
    return g.Wait()
}
```

### Pipeline cancellation

```go
func generator(ctx context.Context) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := 0; i < 100; i++ {
            select {
            case <-ctx.Done(): return
            case out <- i:
            }
        }
    }()
    return out
}
```

### Timers and tickers

```go
t := time.NewTimer(time.Second)
select { case <-t.C: }
t.Stop()

tick := time.NewTicker(time.Second)
defer tick.Stop()
for i := 0; i < 3; i++ { <-tick.C }
```

### Avoiding common pitfalls

- Loop variable capture in goroutines: shadow the variable.
- Buffered channels can still deadlock if buffer fills; always consider receiver progress.
- Always close producer-owned channels; never close a channel you did not create.
- Prefer context over custom done channels for cancellation.

