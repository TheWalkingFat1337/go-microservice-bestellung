FROM golang:latest

WORKDIR /bestellung

COPY . /bestellung

RUN go build -o main

EXPOSE 5678

CMD [ "/bestellung/main" ]