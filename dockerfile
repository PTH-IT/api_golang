FROM golang:1.18.2-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /app/main ./
ENTRYPOINT ["/app/main"]