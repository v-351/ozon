FROM golang:latest

ENV GOPATH=/

COPY cmd/ cmd/
COPY internal/ internal/
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
RUN go build -o url-shortener ./cmd/url-shortener/main.go

CMD [ "./url-shortener" ]