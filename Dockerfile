# Use a imagem oficial do Golang como base
FROM golang:latest

WORKDIR /api-gin

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8080

# Comando para executar o aplicativo
CMD ["./main"]
