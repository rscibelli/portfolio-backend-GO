FROM golang:1.19-bullseye

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY src/main/*.go ./

RUN go build -o /docker-gs-ping

EXPOSE 3000

CMD [ "/docker-gs-ping" ]