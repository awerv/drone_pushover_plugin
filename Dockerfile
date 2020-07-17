FROM golang:1.13.4-alpine as build
RUN mkdir -p /go/src/
WORKDIR /go/src/ 
COPY *.go ./
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-pushover

FROM alpine
COPY --from=build /go/src/drone-pushover /bin/
WORKDIR /bin
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/drone-pushover