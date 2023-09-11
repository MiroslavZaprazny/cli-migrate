FROM golang:1.21.1
WORKDIR /app
COPY ./ /app/
RUN go mod download
