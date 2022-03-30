FROM golang:1.18.0-alpine3.15 as builder

RUN apk update \
  && apk add --no-cache git curl

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /main ./cmd
 
FROM alpine:3.15
 
COPY --from=builder /main .

ENV PORT=${PORT}
ENTRYPOINT ["/main web"]
