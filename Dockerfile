FROM golang:1.13.5-alpine3.10 AS builder
RUN apk --update add git

WORKDIR /go/src/github.com/jimareed/qlik-hackathon-2019-12/
COPY . /go/src/github.com/jimareed/qlik-hackathon-2019-12/
RUN rm -f goapp
RUN go get github.com/gorilla/mux
RUN go build -o ./goapp .

FROM alpine:3.10

EXPOSE 8080
COPY --from=builder /go/src/github.com/jimareed/qlik-hackathon-2019-12/goapp /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/goapp && chmod +x /usr/local/bin/goapp
USER nobody
CMD goapp