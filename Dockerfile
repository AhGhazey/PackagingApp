FROM golang:1.21-alpine3.18 as builder
USER root
RUN apk update && apk add g++ make openssh-client bash git

WORKDIR /go/src/github/ahmedghazey/packaging
#install certificate here in the future

ADD Makefile .
ADD go.mod .
ADD go.sum .

RUN --mount=type=ssh make mod
COPY . .
RUN --mount=type=ssh make build

# Main image
FROM golang:1.21-alpine3.18
COPY --from=builder /go/src/github/ahmedghazey/packaging/packaging /go/src/github/ahmedghazey/packaging/packaging
COPY --from=builder /go/src/github/ahmedghazey/packaging/app.env /app.env

EXPOSE 7070

ENTRYPOINT ["/bin/sh", "-c", "/go/src/github/ahmedghazey/packaging/packaging"]
