# Start from the latest golang base image
# alpine is lightweight Linux distribution, popular in containerized environments
FROM golang:1.21.7-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first; they are less frequently changed than source code, so Docker can cache this layer
COPY go.mod go.sum ./

# Download and verify all dependencies
RUN go mod download 
RUN go mod verify

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Set environment variables (.env)
ENV PORT=8080
ENV LMWN_COVID_ENDPOINT=https://static.wongnai.com/devinterview/covid-cases.json

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]