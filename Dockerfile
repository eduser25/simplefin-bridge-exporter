FROM golang:latest

WORKDIR /go/src/simplefin-bridge-exporter
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/simplefin-bridge-exporter cmd/main.go

FROM gcr.io/distroless/base
COPY --from=0 /go/src/simplefin-bridge-exporter/bin/simplefin-bridge-exporter /simplefin-bridge-exporter 

ENTRYPOINT ["/simplefin-bridge-exporter"]