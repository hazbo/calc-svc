FROM golang:1.14

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0

WORKDIR /go/src/github.com/hazbo/calc-svc
COPY . .
RUN go build \
	-o main \
	-mod vendor \
	-ldflags '-extldflags "-static"' \
	-tags 'osusergo netgo static_build' \
	cmd/calc-server/main.go

# TODO: verify the binary checksum
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
    wget -qO./grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x ./grpc_health_probe


FROM scratch

COPY --from=0 /go/src/github.com/hazbo/calc-svc/main /main
COPY --from=0 /go/src/github.com/hazbo/calc-svc/grpc_health_probe /grpc_health_probe

CMD ["/main"]
