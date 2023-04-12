FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build
RUN go build -o /website ./cmd/web

EXPOSE 8080

# Run
CMD ["/website"]