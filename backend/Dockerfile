#---Build stage---
FROM golang:1.17 AS builder
COPY . /go/src/
WORKDIR /go/src/cmd/app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/service

WORKDIR /go/src/cmd/updatePlans
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/updatePlans

#---Final stage---
FROM alpine:latest
COPY --from=builder /go/bin/service /go/bin/service
COPY --from=builder /go/bin/updatePlans /go/bin/updatePlans
CMD /go/bin/service --port 8090 --host '0.0.0.0'