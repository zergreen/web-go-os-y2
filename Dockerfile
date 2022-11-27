FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN cd src
RUN go build -o main .

CMD ["/app/main"]