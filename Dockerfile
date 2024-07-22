# Stage 1: build the app
FROM golang:latest AS build
WORKDIR /app
COPY ./ ./
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kaudit .

# Stage 2: package the binary
FROM ubuntu:20.04
RUN apt-get update \
    && apt-get install -y ca-certificates \
    && apt-get install -y auditd audispd-plugins \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*


WORKDIR /
RUN mkdir -p /bin
# Copy the binary from the build stage
COPY --from=build /app/kaudit ./bin
ENV PATH="/bin:${PATH}"

