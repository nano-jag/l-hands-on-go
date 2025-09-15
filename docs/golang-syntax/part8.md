## Part 8: Testing, Benchmarking, and Profiling

### Unit tests with `testing`

```go
func TestAdd(t *testing.T) {
    if got := Add(2, 3); got != 5 { t.Fatalf("want 5, got %d", got) }
}
```

### Table-driven tests and subtests

```go
func TestParse(t *testing.T) {
    cases := []struct{name, in string; ok bool}{
        {"empty", "", false},
        {"good", "x=1", true},
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            t.Parallel()
            _, err := Parse(tc.in)
            if (err == nil) != tc.ok { t.Fatalf("ok=%v, err=%v", tc.ok, err) }
        })
    }
}
```

### Test helpers and cleanup

```go
func mustTempDir(t *testing.T) string {
    t.Helper()
    d := t.TempDir()
    return d
}

func TestWithCleanup(t *testing.T) {
    t.Cleanup(func(){ /* runs even on failure */ })
}
```

### Benchmarks

```go
func BenchmarkSum(b *testing.B) {
    data := make([]int, 1000)
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _ = Sum(data)
    }
}
```

Run: `go test -bench=. -benchmem ./...`

### Examples (doc-tested)

```go
func ExampleAdd() {
    fmt.Println(Add(2,3))
    // Output: 5
}
```

### Fuzzing (Go 1.18+)

```go
func FuzzParse(f *testing.F) {
    f.Add("a=1")
    f.Fuzz(func(t *testing.T, s string) {
        _ = Parse(s)
    })
}
```

Run: `go test -fuzz=Fuzz -fuzztime=10s`

### Race detector and coverage

```bash
go test -race ./...
go test -cover -coverprofile=cover.out ./...
go tool cover -html=cover.out
```

### Profiling with pprof

```bash
go test -bench=BenchmarkSum -benchmem -cpuprofile cpu.out -memprofile mem.out ./pkg
go tool pprof cpu.out
```

For servers, expose `net/http/pprof`:

```go
import _ "net/http/pprof"
// http.ListenAndServe("localhost:6060", nil)
```

### Golden files and testdata

- Place fixtures under `testdata/`; `go list` and tools ignore it.
- Use `cmp.Diff` or `go-cmp` to compare structures with readable diffs.

### Determinism and flaky tests

- Seed randomness with a fixed seed or inject RNG.
- Add timeouts to network/external tests; use context.
- Avoid relying on timing; use channels and synchronization.

