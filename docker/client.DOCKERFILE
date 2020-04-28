FROM golang:alpine as builder


RUN apk update \
    && apk add --no-cache git

COPY ./sources /go/src/github.com/jvikstrom/devops-demo/sources
WORKDIR /go/src/github.com/jvikstrom/devops-demo/sources/

# Get dependencies
RUN go get ./...
# ./...
# build the binary
RUN go build


##########################
# Step 2
##########################

FROM golang:alpine

COPY --from=builder /go/src/github.com/jvikstrom/devops-demo/sources/sources /go/bin/client
CMD /go/bin/client -server=hello-server
