FROM golang:1.24

# Install air (live reload)
RUN go install github.com/air-verse/air@latest

# Install Delve debugger
RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app

# Copy go.mod dan go.sum dulu buat cache dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy semua kode
COPY . .

# Set PATH supaya air dan dlv bisa diakses
ENV PATH=$PATH:/go/bin

# Default command: jalankan air untuk live reload
CMD ["air"]
