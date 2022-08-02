# BUILD
FROM golang:1.19-rc-bullseye as build

RUN mkdir app
WORKDIR /app
COPY . .
EXPOSE 9999
#RUN apt update -yq
#RUN apt install -y postgresql-client

#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /http-service .
CMD [ "go","run","main.go" ]
