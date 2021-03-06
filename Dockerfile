FROM golang:1.13 as builder

# setup the working directory
WORKDIR /go/src/app

# dependencies
RUN go get github.com/rs/cors

# add source code
ADD src src

# build the source
RUN go build src/*.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main src/*.go

# use a minimal alpine image
FROM alpine:3.7

# set working directory
WORKDIR /root

# copy the binary from builder
COPY --from=builder /go/src/app/main .
RUN chmod 755 main && chmod 655 . && apk add --no-cache tzdata

# run the binary
CMD ["./main"]
