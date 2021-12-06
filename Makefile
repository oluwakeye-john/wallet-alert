dev:
	APP_ENV=development nodemon --exec go run server.go --signal SIGTERM

generate:
	go generate ./...