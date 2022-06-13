FROM golang:1.18-alpine
WORKDIR /app

#RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o rest-api cmd/main.go

##https://github.com/ufoscout/docker-compose-wait
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait


CMD ["./rest-api"]
EXPOSE 80
