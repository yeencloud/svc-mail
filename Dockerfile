FROM golang:1.24-bullseye AS builder

# Install required packages for runtime using APT
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Cache and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Set build-time arguments (default values can be overridden)
ARG TARGETOS=linux
ARG TARGETARCH=amd64

# Build the Go application with CGO disabled for portability
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o svc cmd/main.go

ARG GITHUB_SHA
ENV GITHUB_SHA=${GITHUB_SHA}

ARG GITHUB_REPOSITORY
ENV GITHUB_REPOSITORY=${GITHUB_REPOSITORY}

ARG GITHUB_REPOSITORY_URL
ENV GITHUB_REPOSITORY_URL=${GITHUB_REPOSITORY_URL}

ARG APP_VERSION
ENV APP_VERSION=${APP_VERSION}

ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}

# Run the application
ENTRYPOINT ["/app/svc"]