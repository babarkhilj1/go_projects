# Multistage Docker build to reduce the size of the final Docker image.
# Stage 1: Build the Go application
FROM golang:alpine as builder

# Create a directory for the build process.
RUN mkdir /build

# Copy all project files into the build directory.
ADD . /build/

# Set the working directory to /build.
WORKDIR /build

# Build the Go application and output the binary as 'main'.
RUN go build -o main .

# Stage 2: Deploy the built application
FROM alpine

# Create a non-root user for running the application (security best practice).
RUN adduser -S -D -H -h /app appuser

# Switch to the non-root user.
USER appuser

# Copy all project files to the /app directory in the container.
COPY . /app

# Copy the built binary from the builder stage to the /app directory.
COPY --from=builder /build/main /app/

# Set the working directory to /app.
WORKDIR /app

# Expose port 3000 for the application.
EXPOSE 3000

# Run the application binary as the container's entry point.
CMD ["./main"]
