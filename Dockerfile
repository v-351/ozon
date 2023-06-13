FROM golang:latest as builder
WORKDIR /build
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener cmd/url-shortener/main.go

FROM scratch
COPY --from=builder build/url-shortener url-shortener
CMD ["./url-shortener"]