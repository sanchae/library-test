FROM golang:1.23.2 as builder
COPY go.mod go.sum /go/src/github.com/sanchae/library-test/
WORKDIR /go/src/github.com/sanchae/library-test
RUN go mod download
COPY . /go/src/github.com/sanchae/library-test
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/library-test github.com/sanchae/library-test

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/sanchae/library-test/build/library-test /usr/bin/library-test
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/library-test"]