FROM golang:alpine3.16 AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /app/server cmd/server/main.go

FROM gcr.io/distroless/base AS runtime
COPY --from=builder /app/server /
CMD ["/server"]