# Start from a base Go image
FROM golang:1.18-alpine

# Set the working directory
WORKDIR /app

# Copy the Go files into the container
COPY . /app/

RUN go mod download
RUN go build -o ./playGo main.go

ENTRYPOINT [ "./playGo" ]
