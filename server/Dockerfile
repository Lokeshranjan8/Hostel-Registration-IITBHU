FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod init server || true && go get github.com/gorilla/mux && go mod tidy

EXPOSE 8080

CMD ["go","run","main.go"]