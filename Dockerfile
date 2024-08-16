# Use an official Go runtime as a parent image
FROM golang:1.21-alpine

# Install necessary packages
RUN apk add --no-cache git

# Set the working directory
WORKDIR /letter

# Clone the GitHub repository containing the Go program
RUN git clone https://github.com/ruvido/last-letter.git .

# Build the Go binary
RUN go build -o letter
RUN chmod +x letter
RUN cp -rp letter /usr/local/bin

# Set the binary as the entrypoint
#ENTRYPOINT ["letter"]
