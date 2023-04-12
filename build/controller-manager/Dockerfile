FROM --platform=$BUILDPLATFORM golang:1.20-alpine as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

ENV GOPROXY=goproxy.io,direct

RUN go mod download

COPY cmd/controller-manager cmd/controller-manager
COPY pkg/ pkg/

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -o /out/kuberiter-controller-manager ./cmd/controller-manager/main.go


FROM alpine:3.14
WORKDIR /
COPY --from=builder /out/kuberiter-controller-manager .
USER 65532:65532

ENTRYPOINT ["./kuberiter-controller-manager"]
