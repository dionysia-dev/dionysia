FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GOCACHE=/root/.cache/build

# build app
RUN --mount=type=cache,target="/root/.cache/build" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# gpac
FROM gpac/ubuntu

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]
