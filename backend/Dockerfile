FROM golang:1.19-alpine AS build
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM build as prod
COPY . .
RUN go build -o /dion
EXPOSE 8000
CMD [ "/dion" ]

FROM build as development
RUN go get github.com/cosmtrek/air
RUN go install github.com/cosmtrek/air
EXPOSE 8000
CMD ["air"]
