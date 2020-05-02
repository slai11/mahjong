FROM golang:1.14.2-stretch
# Copy project into docker instance
COPY . /app
WORKDIR /app
RUN mkdir -p /go/src/mahjong/
RUN cp *.go /go/src/mahjong/

# Get dependencies
RUN go mod download && go build

# Expose 9091 port
EXPOSE 9091/tcp

# Set entrypoint command
CMD ./mahjong
