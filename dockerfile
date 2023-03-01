FROM golang:1.18.2-alpine
WORKDIR /app
RUN go mod tidy
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /out/main ./
ENTRYPOINT ["/out/main"]