FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

RUN go get -d -v ./...

# Expose port 3000 to the outside world
EXPOSE 3000

CMD ["go", "run", "cmd/main.go"]