# run the platform using docker compose
run:
	docker compose up

# stop the platform
stop:
	docker compose down

# run test suite
test:
	go test ./...

# run test suite with integration tests
test-all:
	go test -tags=integration ./...

# lint the code
lint:
	golangci-lint run -v ./...

# generate swagger docs
docs:
	swag init -g internal/api/api.go -o docs

# ingest a real video in loop
ingest $uuid:
		ffmpeg -re -stream_loop -1 -i $(pwd)/files/big_buck_bunny.mp4 \
		-vf "drawtext=fontfile=files/OpenSans-Bold.ttf:fontsize=24:fontcolor=white:x=10:y=10:box=1:boxcolor=black@0.5:boxborderw=5:text='%{pts\:hms}'" \
		-c:v libx264 -preset ultrafast -tune zerolatency -profile:v high \
		-f flv rtmp://localhost:1935/$uuid
