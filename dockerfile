FROM golang:1.17 as base

WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /build

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/build/main"]