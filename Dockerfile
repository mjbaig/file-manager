FROM golang:1.16

WORKDIR /go/src/app
COPY . .

ENV PORT=8080

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["file-manager"]