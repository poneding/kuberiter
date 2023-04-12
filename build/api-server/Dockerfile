FROM --platform=$BUILDPLATFORM golang:1.20-alpine as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

ENV GOPROXY=goproxy.io,direct

RUN go mod download

COPY cmd/api-server cmd/api-server
COPY pkg/ pkg/

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -o /out/kuberiter-api-server ./cmd/api-server/main.go


FROM alpine:3.14
WORKDIR /
COPY --from=builder /out/kuberiter-api-server .
USER 65532:65532

ENTRYPOINT ["./kuberiter-api-server"]
