FROM golang:1.19

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o run

EXPOSE 8090

CMD ["/app/run"]