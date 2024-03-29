FROM golang:1.19-rc-bullseye

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY src/main/*.go ./

RUN go mod tidy

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]