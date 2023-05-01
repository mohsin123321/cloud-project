FROM alpine:3.17

WORKDIR /app

# Copy the app binary
COPY goapp .

# Copy the app configuration
COPY config.json .

ENTRYPOINT ["/app/goapp"]