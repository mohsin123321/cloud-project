FROM alpine:3.17

WORKDIR /app

# Copy the app binary
COPY go-app .

# Copy the app configuration
COPY config.json .

ENTRYPOINT ["/app/go-app"]