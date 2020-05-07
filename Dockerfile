FROM golang:1.14.2-stretch AS build-env
# Copy project into docker instance
COPY . /app
WORKDIR /app
RUN mkdir -p /go/src/mahjong/

# Get dependencies
RUN go mod download && go build

FROM alpine
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=build-env /app/mahjong /app

EXPOSE 80/tcp

CMD ./mahjong
