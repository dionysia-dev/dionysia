build-media-server:
	docker build -f media-server/Dockerfile -t media-server .

media-server:
	docker run --rm --network host -it \
		-v $(PWD)/media-server/mediamtx.yml:/mediamtx.yml \
		-e MTX_PROTOCOLS=tcp \
		-e MTX_WEBRTCADDITIONALHOSTS=192.168.x.x \
		media-server
run:
	docker compose up

stop:
	docker compose down

test:
	go test -v ./...

lint:
	golangci-lint run -v ./...

docs:
	swag init -g internal/api/api.go -o docs
