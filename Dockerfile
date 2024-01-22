# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . .

# Install any dependencies
RUN go get -u github.com/cosmtrek/air

# Expose port 8080 to the outside world
EXPOSE 8080

# Run Air for hot reloading
CMD ["air"]