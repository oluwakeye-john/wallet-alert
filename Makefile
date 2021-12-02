dev:
	nodemon --exec go run server.go --signal SIGTERM

generate:
	go generate ./...