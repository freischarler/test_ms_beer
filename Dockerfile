# Create build stage based on buster image
# Create working directory under /app
FROM golang:1.16-buster AS builder

# Add Maintainer info
LABEL maintainer="Martin Paz <martin.paz@live.com.ar>"

# Set Environment Variables
ENV IP "0.0.0.0"

# Copy over all go config (go.mod, go.sum etc.)
WORKDIR /hex
# Install any required modules
COPY . .
COPY .env .   
# Copy over Go source code
RUN go mod download
# Make sure to expose the port the HTTP server is using
RUN go build -o /out
# Run the app binary when we run the container
EXPOSE 9000
ENTRYPOINT ["/out"]

#docker build -t hex .
#docker run -d -p 9000:9000 -t hello                  
#curl http://localhost:9000/beers                     #test