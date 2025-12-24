FROM golang:1.24-alpine AS builder
RUN apk --update add build-base

WORKDIR /src/app
# Adding go.mod and go.sum and downloading dependencies first
# This is done to leverage Docker layer caching
ADD go.* .
RUN go mod download

# Downloading the tailwind binary, musl because this is an Alpine image.
# This is done first to leverage Docker layer caching
RUN go tool tailo download -v v4.0.6 --musl

ADD . .

# Generating the Tailwind CSS styles with the tailwind binary previously downloaded.
RUN go tool tailo --i internal/system/assets/tailwind.css -o internal/system/assets/application.css

# Building the app with necessary tags - FORCE STATIC COMPILATION
RUN CGO_ENABLED=0 go build -tags osusergo,netgo -ldflags="-s -w" -o bin/app ./cmd/app

# Use Debian instead of Alpine for LibreOffice compatibility
FROM debian:bookworm-slim

# Install LibreOffice with Thai font support
RUN apt-get update && apt-get install -y --no-install-recommends \
    libreoffice \
    fonts-thai-tlwg \
    fonts-noto \
    ca-certificates \
    tzdata \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /bin/

# Copying binaries to /bin from the builder stage
COPY --from=builder /src/app/bin/app .

# Specifying the shell to use
SHELL ["/bin/bash", "-c"]
CMD ["./app"]