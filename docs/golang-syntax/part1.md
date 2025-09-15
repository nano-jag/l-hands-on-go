## Part 1: Basic Syntax, Variables, and Type System Nuances

### Variable declarations and initialization

```go
var name string            // zero value
var age = 25               // type inference
name := "John"            // short declaration (most common)

var (
    firstName string
    lastName  string
    n         int
)

var a, b, c int = 1, 2, 3
x, y := 10, 20
```

### Zero values

```go
var i int            // 0
var f float64        // 0.0
var b bool           // false
var s string         // ""
var p *int           // nil
var slice []int      // nil
var m map[string]int // nil
var ch chan int      // nil
```

### Type conversion

```go
var i int = 42
var f float64 = float64(i) // explicit only
var s string = string(i)   // Unicode code point, not "42"

// For numeric <-> string, use strconv
// s := strconv.Itoa(i) // "42"
```

### Constants and iota

```go
const Pi = 3.14159
const (
    StatusOK       = 200
    StatusNotFound = 404
)

const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
)

const (
    _  = iota
    KB = 1 << (10 * iota)
    MB
    GB
)

const (
    ReadPermission = 1 << iota
    WritePermission
    ExecutePermission
)
```

### Type aliases vs type definitions

```go
type UserID = int  // alias; same type
type UserID2 int   // new type; requires conversion
```

### Custom types and methods

```go
type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit { return Fahrenheit(c*9/5 + 32) }
```

### Pointers

```go
var x int = 42
var p *int = &x
*p = 100 // mutates x
```

### Arrays vs slices

```go
arr := [3]int{1,2,3} // value type (copied on assign)
sl  := []int{1,2,3}  // reference to backing array

sl = append(sl, 4, 5)
sub := sl[1:3]      // shares backing array
subCap := sl[1:3:3] // capacity-limited slice

dst := make([]int, len(sl))
copy(dst, sl)       // independent copy
```

### Maps

```go
m := map[string]int{"apple":5}
if v, ok := m["apple"]; ok { _ = v }
delete(m, "apple")
```

### Control structures

```go
if err := f(); err != nil { return err }

switch day := time.Now().Weekday(); day {
case time.Monday:
    // ...
default:
    // ...
}

for i := 0; i < 3; i++ { /* ... */ }
for cond() { /* ... */ }
for { break }
```

### Strings, bytes, runes

```go
s := "Hello, 世界"
_ = len(s)          // bytes
_ = len([]rune(s))  // runes
for i, r := range s { _ = i; _ = r }
```

