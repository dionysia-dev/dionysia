FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# build app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# gpac
FROM gpac/ubuntu

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]
