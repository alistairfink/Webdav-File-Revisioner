FROM golang:1.16.3 as builder
LABEL maintainer="AlistairFink <alistairfink@gmail.com>"

WORKDIR /go/src/github.com/alistairfink/Webdav-File-Revisioner
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/Webdav-File-Revisioner .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/Webdav-File-Revisioner .
COPY --from=builder /go/src/github.com/alistairfink/Webdav-File-Revisioner/config.json .

CMD ["./Webdav-File-Revisioner"]