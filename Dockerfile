FROM golang:1.19-bullseye

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY src/main/*.go ./

RUN go build -o /docker-gs-ping

COPY ../../persistentFiles/weddingSongs weddingSongs

EXPOSE 8080

CMD [ "/docker-gs-ping" ]