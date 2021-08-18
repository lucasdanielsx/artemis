FROM golang:1.16-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/artemis

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
#RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/artemis/ ./...

COPY config/.env-example ./out/artemis/config/.env

FROM alpine:3.9

RUN apk add ca-certificates

COPY --from=build_base /tmp/artemis/out/artemis /app/artemis

EXPOSE 8080

WORKDIR /app/artemis

CMD ["/app/artemis/main"]