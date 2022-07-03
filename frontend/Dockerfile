#---Build front---
FROM node:17 AS front
WORKDIR /app
COPY . .
ARG REACT_APP_BACKEND_URL
ENV REACT_APP_BACKEND $REACT_APP_BACKEND_URL
RUN npm install --force && npm run build

#---Build server---
FROM golang:1.17 AS back
COPY ./simple-server /go/src/
WORKDIR /go/src/cmd/app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/service

FROM alpine:latest
WORKDIR /go/bin
COPY --from=front /app/build /go/bin/build
COPY --from=back /go/bin/service /go/bin/service
CMD /go/bin/service --port 8090 --host '0.0.0.0'
