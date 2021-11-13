# dev, builder
FROM golang:1.16 AS golang
WORKDIR /work/yatter-backend-go

# dev
FROM --platform=linux/amd64 golang:1.16 as dev
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go get -v github.com/rubenv/sql-migrate/...
# builder
FROM golang AS builder
COPY ./ ./
RUN make prepare build-linux

# release
FROM alpine AS app
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=builder /work/yatter-backend-go/build/yatter-backend-go-linux-amd64 /usr/local/bin/yatter-backend-go
EXPOSE 8080
ENTRYPOINT ["yatter-backend-go"]
