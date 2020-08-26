# with golang version should be used? (i saw somwthing about alpine3.9 in tutorials)
FROM golang
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build ./cmd/simple-service/main.go
CMD ["/app/main"]