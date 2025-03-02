FROM golang:1.24

# RUN apk update && apk upgrade && \
#     apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

EXPOSE 9000

CMD ["./main"]
