FROM golang:1.18.2-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=linux  go build  -o /out/main ./
ENTRYPOINT ["/out/main"]