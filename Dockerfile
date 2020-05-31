FROM golang:1.14.3-alpine3.11
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
ENTRYPOINT /app/main
EXPOSE 5000