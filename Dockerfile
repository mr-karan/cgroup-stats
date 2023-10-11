# Build stage
FROM golang:1.21 AS build
WORKDIR /app
COPY . .
RUN go build -o /app/cgroups.bin examples/main.go

# Run stage 
FROM ubuntu:latest  
WORKDIR /app
COPY --from=build /app/cgroups.bin /app/
CMD ["/app/cgroups.bin"]  
