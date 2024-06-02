FROM golang:alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download & go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/url-checker/main.go
FROM scratch
COPY --from=builder /src/app .
COPY --from=builder /src/config.json .
CMD ["/app"]