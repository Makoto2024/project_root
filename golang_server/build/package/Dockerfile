# Build the Go binary with bazel
FROM golang:1.24 AS bazel_builder
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl && \
    curl -fSL "https://github.com/bazelbuild/bazelisk/releases/download/v1.26.0/bazelisk-linux-amd64" -o /usr/local/bin/bazelisk && \
    chmod +x /usr/local/bin/bazelisk && \
    ln -s /usr/local/bin/bazelisk /usr/local/bin/bazel

WORKDIR /app
COPY . .
RUN bazel build //golang_server/cmd/server

# Stage 2: Create the final minimal image
FROM debian
COPY --from=bazel_builder /app/bazel-bin/golang_server/cmd/server/server_/server /server
EXPOSE 8080
ENTRYPOINT [ "/server" ]
