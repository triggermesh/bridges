FROM golang:1.15-buster AS builder

ENV CGO_ENABLED 0

WORKDIR /project
COPY . ./

RUN go build

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /project/translator /

ENTRYPOINT ["/translator"]
