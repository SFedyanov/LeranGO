FROM golang:1.19.4

WORKDIR /usr/src/app

COPY . .

RUN go build hello.go

CMD ["./hello"]

# docker build -t my-golang-app .
# docker run -it --rm --name my-running-app my-golang-app
