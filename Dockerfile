# Use the latest Alpine as the base image
FROM alpine:latest

# Install necessary dependencies for building Go applications
RUN apk add --no-cache ca-certificates git

# Set the latest Go version as an environment variable
ENV GO_VERSION=1.21.0

# Download and install the latest Go binary
RUN wget -q https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

# Add Go binaries to the system path
ENV PATH="/usr/local/go/bin:${PATH}"

# Set the working directory
WORKDIR /go/src/github.com/lindsaygelle/slice

# Copy the application source code into the container
COPY . .

# Download and cache the Go module dependencies
RUN go mod download
