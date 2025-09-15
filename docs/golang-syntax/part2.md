## Part 2: Functions, Methods, and Interfaces

### Function signatures and returns

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums { total += n }
    return total
}
```

Notes:
- Parameters are passed by value. Use pointers/reference types to mutate.
- Variadics are slices; only one, must be last.
- Named returns create variables at function entry; use sparingly.

### Defer

```go
func withFile(p string) error {
    f, err := os.Open(p)
    if err != nil { return err }
    defer f.Close()
    return nil
}
```

### Methods and receivers

```go
type Counter struct{ value int }
func (c *Counter) Inc() { c.value++ }
func (c Counter) Value() int { return c.value }
```

Guidelines:
- Pointer receivers for mutation/large structs/interface satisfaction.
- Value receivers fine for small, immutable-like types.

### Interfaces and implementation

```go
type Reader interface{ Read(p []byte) (int, error) }
type MyBuf struct{}
func (MyBuf) Read(p []byte) (int, error) { return 0, io.EOF }
```

Notes:
- Implicit, structural satisfaction. Keep interfaces small; define in consumer.

### Nil and interfaces

```go
var e error               // nil
var p *os.PathError = nil // typed nil
var err error = p         // non-nil interface
_ = (err == nil)          // false
```

### Function values and closures

```go
func makeFuncs() []func() int {
    fs := make([]func() int, 0, 3)
    for i := 0; i < 3; i++ {
        i := i // shadow to capture value
        fs = append(fs, func() int { return i })
    }
    return fs
}
```

### Init functions

```go
func init() { /* light registration only */ }
```

### Function options pattern

```go
type Server struct{ addr string; tls bool }
type Option func(*Server)

func WithTLS(b bool) Option { return func(s *Server) { s.tls = b } }
func WithAddr(a string) Option { return func(s *Server) { s.addr = a } }

func NewServer(opts ...Option) *Server {
    s := &Server{addr: ":8080"}
    for _, opt := range opts { opt(s) }
    return s
}
```

### Method sets and interfaces

```go
type S struct{}
func (S) A() {}
func (*S) B() {}

var _ interface{ A() } = S{}    // ok
// var _ interface{ B() } = S{}  // not ok
var _ interface{ B() } = &S{}   // ok
```

### Generics (Go 1.18+)

```go
type Number interface{ ~int | ~int64 | ~float64 }
func Sum[T Number](v []T) T { var t T; for _, x := range v { t += x }; return t }
```

