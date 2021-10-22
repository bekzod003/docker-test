FROM golang:alpine
WORKDIR /go/delivery

COPY . ./
RUN go build -o main .
CMD ["./main"]