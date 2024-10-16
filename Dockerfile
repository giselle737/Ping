# Use the official Golang image as a base
FROM golang:1.23.2

# Set the working directory inside the container
WORKDIR /app

# Copy the stock_ticker directory to the container
COPY ./stock_ticker /app/stock_ticker

# Set the working directory where main.go is located
WORKDIR /app/stock_ticker

# Download dependencies and tidy the modules
RUN go mod tidy

# Expose the port that the application will run on
EXPOSE 8080

# Run the application directly
CMD ["go", "run", "main.go"]
