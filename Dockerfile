FROM golang:latest as build

RUN mkdir app
WORKDIR /app
COPY . .

RUN go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo && go get github.com/onsi/gomega/...

EXPOSE 9999

CMD [ "go","run","main.go" ]
