FROM golang:1.14-stretch AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download


COPY . .
RUN go install -installsuffix cgo .


FROM registry.access.redhat.com/ubi8/ubi-minimal
LABEL name "Triggermesh Zenformation"
LABEL vendor "Triggermesh"
LABEL version "v0.1.0"
LABEL release "1"
LABEL summary "Triggermesh Zenformation"
LABEL description "This is the Triggermesh Zenformation"


COPY --from=builder /go/bin/main /go/bin/main

ENTRYPOINT ["/go/bin/main"]
