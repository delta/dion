FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o /dion

EXPOSE 8000

CMD [ "/dion" ]
