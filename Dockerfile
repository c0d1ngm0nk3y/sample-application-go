FROM golang:latest
RUN mkdir /app
RUN mkdir -p /go/src
RUN export GOPATH=$GOPATH:/app:/go/src
ADD . /app
WORKDIR /app
RUN go get -d -v ./...
RUN go build -o simple-app .
RUN ls simple-app
EXPOSE 8080
CMD ["/app/simple-app"]
