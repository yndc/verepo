
#--------------------------------
# Stage 1 - Builder
#--------------------------------
FROM golang:1.18rc1-bullseye AS builder
WORKDIR /app
ARG APP_ID
ENV APP_ID=${APP_ID}

# Get dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download 

# Copy the app
COPY api ./api
COPY cmd ./cmd
COPY pkg ./pkg
COPY internal ./internal

# Build the app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags static -ldflags="-w -s" -o ./build/${APP_ID} ./cmd/${APP_ID}/.

#--------------------------------
# Stage 2 - Deployment container
#--------------------------------
FROM scratch
ARG APP_ID
ENV APP_ID=${APP_ID}

# Install certificates
ADD ca-certificates.crt /etc/ssl/certs/

# Copy the compiled app
COPY --from=builder /app/build/${APP_ID} /app

# Run the binary
ENTRYPOINT ["/app"]
