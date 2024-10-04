# Use the official Golang image as the base image with Debian
FROM golang:1.23-bullseye AS builder

# Set the working directory
ARG APP_RELATIVE_PATH
ARG service
ENV SERVICE_NAME $service

COPY . /data/app
WORKDIR /data/app

RUN rm -rf /data/app/bin/
RUN go mod vendor && go build -ldflags="-s -w" -o ./bin/app .
#RUN mv config /data/app/bin/

# Use a minimal Debian base image for the final stage
FROM debian:bullseye-slim

# Install necessary packages
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy the built binary and other required files from the builder stage

WORKDIR /data/app
COPY --from=builder /data/app/bin /data/app
#COPY --from=builder /data/app/public /data/app/public

ENTRYPOINT ["./app"]
