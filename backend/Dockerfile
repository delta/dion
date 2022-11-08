:FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /dion

EXPOSE 8000

CMD [ "/dion" ]

