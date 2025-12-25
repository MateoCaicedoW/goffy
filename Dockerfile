FROM golang:1.24-alpine AS builder
RUN apk --update add build-base

WORKDIR /src/app
ADD go.* .
RUN go mod download

RUN go tool tailo download -v v4.0.6 --musl

ADD . .

RUN go tool tailo --i internal/system/assets/tailwind.css -o internal/system/assets/application.css

# Building the app with necessary tags - FORCE STATIC COMPILATION
RUN CGO_ENABLED=0 go build -tags osusergo,netgo -ldflags="-s -w" -o bin/app ./cmd/app

FROM debian:bookworm-slim

# Install LibreOffice, Python, and pdf2docx with comprehensive font support
RUN apt-get update && apt-get install -y --no-install-recommends \
    libreoffice \
    python3 \
    python3-pip \
    python3-venv \
    fonts-thai-tlwg \
    fonts-noto \
    fonts-noto-cjk \
    fonts-noto-color-emoji \
    fonts-liberation \
    fonts-dejavu-core \
    fonts-dejavu-extra \
    fonts-crosextra-carlito \
    fonts-crosextra-caladea \
    ca-certificates \
    tzdata \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Install pdf2docx globally (no virtual environment needed in container)
RUN pip3 install --no-cache-dir pdf2docx --break-system-packages

WORKDIR /bin/

COPY --from=builder /src/app/bin/app .

SHELL ["/bin/bash", "-c"]
CMD ["./app"]