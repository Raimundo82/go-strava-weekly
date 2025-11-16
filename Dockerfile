FROM golang:1.25.4 AS builder

WORKDIR /app

# Download dependencies separately for caching
COPY go.mod ./
RUN go mod download

# Copy source
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o strava-weekly

# ---------- Final Image ----------
FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/strava-weekly /strava-weekly

USER nonroot:nonroot

ENTRYPOINT ["/strava-weekly"]
