FROM golang:1.19.3-alpine AS build
RUN apk add build-base
WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY ./cmd ./cmd
COPY ./pkg ./pkg
RUN go mod download

RUN go build -o fetch ./cmd/fetch/main.go

FROM alpine:latest
WORKDIR /data
CMD ["sh"]

COPY --from=build /app/fetch /app/fetch
RUN ln -s /app/fetch /bin/fetch