# run the platform using docker compose
run:
	docker compose up

# stop the platform
stop:
	docker compose down

# run test suite
test:
	go test ./...

# lint the code
lint:
	golangci-lint run -v ./...

# generate swagger docs
docs:
	swag init -g internal/api/api.go -o docs
