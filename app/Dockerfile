FROM golang:1.8

# Expose the application on port 8080
EXPOSE 8080

COPY  ./primeServer /go/src/primeServer
RUN go build -o primeServer /go/src/primeServer/main.go
RUN go install primeServer

CMD ["primeServer"]