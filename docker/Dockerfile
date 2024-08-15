# Use Alpine as the base image
FROM alpine:latest

# Install necessary dependencies. Adjust as needed for your executable.
RUN apk add --no-cache \
    bash  # Install other dependencies if needed

# Set the working directory inside the container
WORKDIR /app

# Copy all files from the current directory to /app in the container
COPY . /app/

# Check if there are letters to send; abort with an error if not
RUN if [ -z "$(ls *.letter 2>/dev/null)" ]; then \
        echo "Error: No letters found to send"; \
        exit 1; \
    fi

# Make the executable accessible
RUN chmod +x ./letter

# Define the entry point for the container
ENTRYPOINT ["./letter", "send"]
