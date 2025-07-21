FROM golang:1.20 AS builder

WORKDIR /app

COPY main.go .
RUN go build -o scanner .

FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

# Install necessary packages
RUN apt-get update && \
    apt-get install -y \
    clamav \
    rkhunter \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Update ClamAV database
RUN freshclam

# Copy the compiled Go program from the builder stage
COPY --from=builder /app/scanner /usr/local/bin/scanner

ENTRYPOINT ["/usr/local/bin/scanner"]

