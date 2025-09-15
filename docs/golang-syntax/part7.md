## Part 7: Package Management and Module System

### Go modules basics

```bash
go mod init example.com/my/mod
go get github.com/acme/lib@v1.2.3
go build ./...
go test ./...
```

Files:
- `go.mod`: module path, Go version, requirements, replace directives.
- `go.sum`: checksums for reproducible builds.

### Versioning and imports

- Semantic import versioning: v2+ requires `/vN` major suffix in module path and import.
- Use tagged releases. Avoid pseudo-versions in libraries.

```go
import (
    lib "github.com/acme/lib/v2"
)
```

### Replace and retract

```go
// go.mod
replace example.com/internal => ../internal

retract [v1.2.0, v1.2.3] // bad releases
```

### Vendoring

```bash
go mod vendor
go build -mod=vendor ./...
```

Use vendoring to snapshot deps in regulated or hermetic environments.

### `go get`, `go work`, and multi-module repos

```bash
go work init ./cmd/app ./lib
go work use ./tools
```

Use workspaces to develop across multiple modules without replace hacks.

### Build tags and constraints

```go
//go:build linux && amd64
// +build linux,amd64 // legacy
```

Place build tags at the top of files. Prefer the `//go:build` form.

Common constraints: `linux`, `darwin`, `windows`, `arm64`, `amd64`, `race`, `cgo`.

### Tooling flags

```bash
go build -ldflags "-s -w" -trimpath ./cmd/app
go test -race -cover ./...
go test -run TestName -v ./pkg/...
go list -m -u all         # check for updates
go mod tidy               # prune unneeded deps
```

### Private modules

- Set `GOPRIVATE=example.com/*` to skip checksum DB and proxy.
- Configure auth via SSH/HTTPS; for GitHub, use tokens with `git`.

### Reproducibility tips

- Pin tools with separate `tools.go` and module or use `go install tool@version`.
- Avoid `replace` in libraries; ok in apps.
- CI: cache module download dir and build cache for speed.

