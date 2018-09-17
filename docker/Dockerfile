FROM golang:latest 
RUN mkdir /app 
WORKDIR /app 
RUN go get -v github.com/c0d1ngm0nk3y/sample-application-go/...
RUN go build -o sample-app github.com/c0d1ngm0nk3y/sample-application-go
RUN ls sample-app
EXPOSE 8080
CMD ["/app/sample-app"]
