FROM golang:1.23.4 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
ENV PATH="/go/bin:$PATH"
RUN swag init

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM gcr.io/distroless/static

WORKDIR /
COPY --from=builder /app/main .
EXPOSE 80
ENTRYPOINT ["/main"]

