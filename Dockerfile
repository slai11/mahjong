FROM golang:1.14.2-alpine

RUN apk --no-cache add ca-certificates
# Copy project into docker instance
COPY . /app
WORKDIR /app
RUN mkdir -p /go/src/mahjong/

RUN go mod download && go build

EXPOSE 80/tcp

CMD ./mahjong
