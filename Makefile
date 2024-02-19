include .env

seed:
	go run db/seeds/seed.go

compile:
	protoc --proto_path=./ml --go_out=./ml --go_opt=paths=source_relative --go-grpc_out=./ml --go-grpc_opt=paths=source_relative ml.proto
.PHONY:seed