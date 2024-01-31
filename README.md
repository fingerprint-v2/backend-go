# backend-go

# Running

- `air`
- `air -c .air.nobuild.toml` (Faster reload)

# Issues

- When running `go generate` command, I got error `cmd/go: missing sum after updating a different package`.
  - I fixed this with `go env -w GOFLAGS=-mod=mod` https://github.com/golang/go/issues/44129

# Makefile

- Windows `choco install make` https://stackoverflow.com/a/54086635
