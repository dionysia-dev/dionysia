stream-server:
	docker run --rm -it \
		-v $(PWD)/mediamtx.yml:/mediamtx.yml \
		-e MTX_PROTOCOLS=tcp \
		-e MTX_WEBRTCADDITIONALHOSTS=192.168.x.x \
		-p 8554:8554 \
		-p 1935:1935 \
		-p 8888:8888 \
		-p 8889:8889 \
		-p 8890:8890/udp \
		-p 8189:8189/udp \
		bluenviron/mediamtx

lint:
	golangci-lint run -v ./...
