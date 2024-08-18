FROM golang:latest

WORKDIR /go/src/simplefin-bridge-exporter
COPY . .
RUN go build -o ./simplefin-bridge-exporter -ldflags "-linkmode external -extldflags -static" -a cmd/main.go

FROM gcr.io/distroless/static
COPY --from=0 /go/src/simplefin-bridge-exporter /simplefin-bridge-exporter 
ENTRYPOINT ["/simplefin-bridge-exporter"]