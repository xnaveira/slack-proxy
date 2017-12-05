FROM golang:latest
ENV GOPATH=/go/src:/app/vendor
RUN mkdir /app
COPY . /app/
WORKDIR /app
RUN ln -s /app/vendor /app/vendor/src
RUN go build -o main .
EXPOSE 8080
CMD ["/app/main"]
