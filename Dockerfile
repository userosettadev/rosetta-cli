FROM golang:1.22 as builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN VERSION=$(git describe --always --tags) && \
    CGO_ENABLED=0 GOOS=linux go build \
    -mod=readonly \
    -ldflags "-s -w -X github.com/userosettadev/rosetta/build.Version=${VERSION}" \
    -o rosetta

FROM alpine:3
WORKDIR /usr/bin
COPY --from=builder /go/src/app/rosetta .
ENTRYPOINT ["/usr/bin/rosetta"]
