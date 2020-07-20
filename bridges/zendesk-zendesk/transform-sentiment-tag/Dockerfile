FROM golang:1.14-stretch AS builder

ENV CGO_ENABLED 0

WORKDIR /project
COPY . ./

RUN go build

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /project/transform-sentiment-tag /

ENTRYPOINT ["/transform-sentiment-tag"]
