FROM golang:1.16-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/artemis

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
#RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/artemis/ ./...

FROM alpine:3.9

RUN apk add ca-certificates

COPY --from=build_base /tmp/artemis/out/artemis /app/artemis

EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/artemis/main"]