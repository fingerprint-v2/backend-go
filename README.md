# Dependencies

- `go install golang.org/x/tools/cmd/stringer@latest`
- `go install github.com/google/wire/cmd/wire@latest`

# Setting up

- `docker volume create fp2_vol_db`
- `docker volume create fp2_vol_minio`
- `docker volume create fp2_vol_pgadmin`
- `docker network create fp2_net`
- `docker-compose --env-file ./.env up -d`

# Running

- `air`
- `air -c .air.nobuild.toml` (Faster reload)

# Issues

- When running `go generate` command, I got error `cmd/go: missing sum after updating a different package`.
  - I fixed this with `go env -w GOFLAGS=-mod=mod` https://github.com/golang/go/issues/44129

# Makefile

- Windows `choco install make` https://stackoverflow.com/a/54086635

# GRPC

## Golang

### Setup

- `choco install protoc`
- `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
- `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

### Compile

- `make compile`

## Python

- Make sure the root is at `./ml`
- See `README.md`
