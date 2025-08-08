FROM golang:1.24
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/inorder cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/inorder .
COPY config.yaml .
COPY logo.txt .
RUN chmod +x /app/inorder
EXPOSE 4000
CMD ["/app/inorder"]
