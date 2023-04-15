# Start by building the application.
FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o libapi cmd/main.go

FROM gcr.io/distroless/static-debian11

COPY --from=builder /app/libapi /

EXPOSE 3000

CMD ["/libapi"]