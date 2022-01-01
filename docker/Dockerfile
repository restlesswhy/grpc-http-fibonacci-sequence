FROM golang:1.17-buster

#ENV TZ=Europe/Moscow
#RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone


WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 5000
EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build cmd/fib/main.go" --command=./main