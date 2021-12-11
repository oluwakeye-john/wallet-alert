dev:
	go run server.go

watch:
	APP_ENV=development nodemon --exec go run server.go --signal SIGTERM

generate:
	go generate ./...