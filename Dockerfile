# Stage 1: build the app
FROM golang:latest AS build
WORKDIR /app
COPY ./ ./
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kaudit .

# Stage 2: package the binary
FROM alpine:latest

WORKDIR /
RUN mkdir -p /bin
# Copy the binary from the build stage
COPY --from=build /app/kaudit ./bin
ENV PATH="/bin:${PATH}"
