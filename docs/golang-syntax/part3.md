## Part 3: Structs, Embedding, and Composition Patterns

### Struct basics and zero values

```go
type Config struct {
    Addr    string
    Timeout time.Duration
    Enabled bool
}

var c Config
c2 := Config{Addr: ":8080", Enabled: true} // keyed literal
c3 := &Config{Timeout: 5 * time.Second}
```

### Export and tags

```go
type User struct {
    ID   int    `json:"id"`
    name string // unexported; not marshaled
}
```

### Nil vs empty for JSON

```go
type Payload struct {
    Items []string          `json:"items,omitempty"`
    Meta  map[string]string `json:"meta,omitempty"`
}

p := Payload{Items: nil} // {"items":null}
p.Items = []string{}     // {"items":[]}
```

### Copy semantics and equality

```go
type A struct{ X int; Y string }
a := A{1, "x"}; b := a; b.X = 2 // a.X remains 1
_ = a == b // comparable

type B struct{ S []int }
// _ = B{} == B{} // not comparable
```

### Constructors and immutability

```go
func NewConfig(opts ...Option) *Config { cfg := &Config{Addr: ":8080"}; for _, o := range opts { o(cfg) }; return cfg }

type Settings struct{ timeout time.Duration }
func NewSettings() Settings { return Settings{timeout: 5 * time.Second} }
func (s Settings) WithTimeout(d time.Duration) Settings { s.timeout = d; return s }
func (s Settings) Timeout() time.Duration { return s.timeout }
```

### Embedding and promotion

```go
type Repo struct {
    *log.Logger // embedded pointer
    db *sql.DB
}

func (r *Repo) Save() error {
    r.Printf("saving...")
    return nil
}
```

### Conflicts and qualification

```go
type A2 struct{}
func (A2) Do() {}
type B2 struct{}
func (B2) Do() {}
type C struct{ A2; B2 }
// c.Do() // ambiguous; use c.A2.Do() or c.B2.Do()
```

### Method sets with embedding

```go
type E struct{}
func (E) A() {}
func (*E) B() {}

type Outer struct{ E }
// B promoted only for *Outer
```

### Concurrency-safe structs

```go
type SafeCounter struct { mu sync.Mutex; n int }
func (s *SafeCounter) Inc() { s.mu.Lock(); s.n++; s.mu.Unlock() }
```

### Optional fields

```go
type Opts struct { Retries *int `json:"retries,omitempty"` }
```

### Delegation pattern

```go
type Clock interface{ Now() time.Time }
type RealClock struct{}
func (RealClock) Now() time.Time { return time.Now() }

type Scheduler struct{ clock Clock }
func NewScheduler(c Clock) *Scheduler { return &Scheduler{clock: c} }
func (s *Scheduler) RunAt(t time.Time) { _ = s.clock.Now() }
```

