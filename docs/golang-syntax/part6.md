## Part 6: Advanced Patterns (generics, reflection, advanced context, API design)

### Generics essentials (Go 1.18+)

```go
// Constraints
type Number interface{ ~int | ~int64 | ~float64 }

func Sum[T Number](vals []T) T { var s T; for _, v := range vals { s += v }; return s }

// comparable: built-in constraint for == and map keys
func IndexOf[T comparable](xs []T, target T) int {
    for i, x := range xs { if x == target { return i } }
    return -1
}
```

Notes:
- Type inference usually elides explicit type args: `Sum([]int{1,2})`.
- Use `~` to match types with identical underlying types (e.g., custom ints).
- Prefer minimal, behavior-focused constraints; avoid `any` unless necessary.

### Generic helpers and containers

```go
// Set with map[T]struct{}
type Set[T comparable] map[T]struct{}
func (s Set[T]) Add(v T){ s[v] = struct{}{} }
func (s Set[T]) Has(v T) bool { _, ok := s[v]; return ok }

// Map over slice
func Map[A,B any](in []A, f func(A) B) []B {
    out := make([]B, len(in))
    for i, a := range in { out[i] = f(a) }
    return out
}
```

### Type constraints via interface embedding

```go
type Signed interface{ ~int | ~int32 | ~int64 }
type Ordered interface{ Signed | ~uint | ~float32 | ~float64 | ~string }

func Min[T Ordered](a, b T) T { if a < b { return a }; return b }
```

### Reflection basics (use sparingly)

```go
func Fields(v any) []string {
    t := reflect.TypeOf(v)
    if t.Kind() == reflect.Ptr { t = t.Elem() }
    if t.Kind() != reflect.Struct { return nil }
    names := make([]string, t.NumField())
    for i := 0; i < t.NumField(); i++ { names[i] = t.Field(i).Name }
    return names
}
```

Guidelines:
- Reflection is slower and bypasses compile-time checks; isolate it.
- Only exported fields are settable via reflection.
- Use `encoding/json` tags and `reflect.StructTag.Get("json")` for tag-driven behavior.

### Reflect: zero values and addressability

```go
rv := reflect.ValueOf(&x).Elem() // addressable
if rv.CanSet() { rv.SetInt(42) }
```

### Advanced context usage

```go
// Custom key type to avoid collisions
type ctxKey struct{ name string }
var keyUser = ctxKey{"user"}

func WithUser(ctx context.Context, u *User) context.Context { return context.WithValue(ctx, keyUser, u) }
func UserFrom(ctx context.Context) (*User, bool) {
    u, ok := ctx.Value(keyUser).(*User); return u, ok
}
```

Best practices:
- Pass `context.Context` as the first parameter; do not store in structs.
- Use values sparingly for request-scoped data (ids, auth, tracing), not for optional params.
- Always call the returned `cancel` from `WithCancel/Timeout/Deadline`.
- Propagate deadlines to downstream calls; respect `ctx.Done()` in long ops.

### API design considerations

- Keep interfaces small and define them on the consumer side.
- Prefer constructors to raw struct literals for non-trivial types.
- Return concrete types, accept interfaces.
- Be explicit about ownership and lifetimes (start/stop, Close methods).
- Expose zero-value-usable types when possible.

### Unsafe and performance caveats (brief)

- `unsafe` breaks guarantees; use only in low-level, well-tested codepaths.
- Prefer allocation-free hot paths; reuse buffers with `sync.Pool` when profiling shows benefit.
- Avoid reflection in hot loops; prefer generics where possible.

