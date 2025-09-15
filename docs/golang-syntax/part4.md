## Part 4: Error Handling and panic/recover

### Idiomatic error values

```go
func load(path string) ([]byte, error) {
    if path == "" {
        return nil, fmt.Errorf("path is empty")
    }
    b, err := os.ReadFile(path)
    if err != nil {
        return nil, err // return the original error
    }
    return b, nil
}
```

Guidelines:
- Return `(T, error)` with zero-value `T` on error.
- Check errors immediately; handle or propagate.
- Keep messages lowercased without trailing punctuation for consistency.

### Sentinel vs typed errors

```go
var ErrNotFound = errors.New("not found") // sentinel

type ErrRateLimited struct { RetryAfter time.Duration }
func (e ErrRateLimited) Error() string { return fmt.Sprintf("rate limited: retry in %s", e.RetryAfter) }
```

Use sentinels for simple, stable conditions; prefer typed errors when carrying context.

### Wrapping and unwrapping

```go
// wrap with %w to retain cause
if err := do(); err != nil {
    return fmt.Errorf("do failed: %w", err)
}

// errors.Is / errors.As
if errors.Is(err, ErrNotFound) { /* handle */ }

var rl ErrRateLimited
if errors.As(err, &rl) { /* use rl.RetryAfter */ }
```

### Attaching context to errors

```go
func fetch(id string) (Item, error) {
    it, err := dbGet(id)
    if err != nil { return Item{}, fmt.Errorf("db get id=%s: %w", id, err) }
    return it, nil
}
```

Prefer adding identifiers/parameters to error paths. Avoid logging and returning the same error (double logs).

### Temporary and retryable errors

```go
type Temporary interface{ Temporary() bool }

func IsTemporary(err error) bool {
    var t Temporary
    return err != nil && errors.As(err, &t) && t.Temporary()
}
```

### Nil-interface pitfall

```go
func bad() error {
    var pe *os.PathError = nil
    return pe // returns non-nil error interface holding (*PathError)(nil)
}
```

Ensure you return `nil` interface when intended.

### Defer and error handling

```go
func writeFile(p string, data []byte) (err error) {
    f, err := os.Create(p)
    if err != nil { return err }
    defer func() {
        cerr := f.Close()
        if err == nil { err = cerr } // prefer close error if no prior error
    }()
    if _, err = f.Write(data); err != nil { return err }
    return nil
}
```

### Panics: when and where

- Panics are for programmer errors or truly unrecoverable states (e.g., invariant violation).
- Do not use panic for expected error paths.
- Recover only at safe boundaries (goroutine top-level, process entry points) to keep the process running and log.

### Recover pattern

```go
func safeGo(fn func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                // log and continue; include stack
                buf := make([]byte, 1<<16)
                n := runtime.Stack(buf, false)
                log.Printf("panic: %v\n%s", r, buf[:n])
            }
        }()
        fn()
    }()
}
```

### Panic vs log.Fatal

- `log.Fatal` exits immediately (os.Exit) and skips defers; prefer returning errors or panicking and recovering at boundaries.

### Multi-error aggregation

```go
type MultiError []error
func (m MultiError) Error() string {
    var b strings.Builder
    for i, e := range m {
        if i > 0 { b.WriteString("; ") }
        b.WriteString(e.Error())
    }
    return b.String()
}
```

### Testing errors

```go
if !errors.Is(err, ErrNotFound) { t.Fatalf("expected not found, got %v", err) }

var rl ErrRateLimited
if !errors.As(err, &rl) || rl.RetryAfter == 0 { t.Fatal("expected rate limit with retry") }
```

