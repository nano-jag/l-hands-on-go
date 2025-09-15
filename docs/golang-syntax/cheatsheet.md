## Go Quick Start Cheatsheet

### Basics

```go
// declare
var n int
name := "go"

// constants
const Pi = 3.14159

// iota
const (
    A = iota; B; C
)

// type alias vs new type
type ID = int
type UserID int
```

### Collections

```go
arr := [3]int{1,2,3}
sl := []int{1,2,3}
sl = append(sl, 4)
sub := sl[1:3]
dst := make([]int, len(sl)); copy(dst, sl)

m := map[string]int{"a":1}
v, ok := m["a"]
for k, v := range m { _ = k; _ = v }
```

### Control flow

```go
if err := f(); err != nil { return err }

switch x {
case 1:
case 2:
default:
}

for i := 0; i < 3; i++ {}
for cond() {}
for { break }
```

### Functions and methods

```go
func sum(xs ...int) int { s := 0; for _, x := range xs { s += x }; return s }

type Counter struct{ n int }
func (c *Counter) Inc(){ c.n++ }
func (c Counter) Val() int { return c.n }
```

### Interfaces

```go
type Reader interface{ Read([]byte) (int, error) }
```

### Errors

```go
var ErrNotFound = errors.New("not found")
if err := do(); err != nil { return fmt.Errorf("do: %w", err) }
if errors.Is(err, ErrNotFound) {}
```

### Concurrency

```go
ch := make(chan int, 1)
go func(){ ch <- 1; close(ch) }()
for v := range ch { _ = v }

ctx, cancel := context.WithTimeout(context.Background(), time.Second); defer cancel()
select { case <-ctx.Done(): case <-time.After(10*time.Millisecond): }
```

### sync

```go
var mu sync.Mutex; mu.Lock(); /* ... */; mu.Unlock()
var wg sync.WaitGroup; wg.Add(1); go func(){ defer wg.Done() }(); wg.Wait()
```

### Generics

```go
type Number interface{ ~int | ~int64 | ~float64 }
func Sum[T Number](xs []T) T { var s T; for _, x := range xs { s += x }; return s }
```

### Testing

```go
func TestX(t *testing.T){ t.Parallel() }
func BenchmarkX(b *testing.B){ for i:=0;i<b.N;i++{} }
```

### Modules

```bash
go mod init example.com/app
go get github.com/acme/lib@v1.2.3
go mod tidy
```

