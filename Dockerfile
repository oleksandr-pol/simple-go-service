# with golang version should be used? (i saw somwthing about alpine3.9 in tutorials)
FROM golang
RUN mkdir /app
ADD . /app
WORKDIR /app/cmd/simple-service
RUN go build
CMD ["/app/cmd/simple-service/simple-service"]