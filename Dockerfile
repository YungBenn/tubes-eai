FROM golang:1.20.5-bullseye

ENV DB_NAME=defaultdb
ENV DB_PASSWORD="AVNS_xRZ2kjuFTTDMs8nnSAZ"
ENV DB_PORT=18649
ENV DB_USER=avnadmin
ENV DB_HOST="pg-38b0df00-rubenadisuryo22-376c.aivencloud.com"
ENV DB_SSLMODE=require
ENV PORT=3000

WORKDIR /app

COPY . .

RUN go mod download

COPY *.go ./

EXPOSE 3000

CMD [ "go", "run", "main.go" ]
