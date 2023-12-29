include .env

seed:
	go run db/seeds/seed.go

.PHONY:seed