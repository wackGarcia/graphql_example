FROM golang:1.13-alpine AS builder
WORKDIR /go/src/graphql_example
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM golang:1.13-alpine
COPY --from=builder /go/src/graphql_example/.env .
COPY --from=builder /go/src/graphql_example/app .
CMD ["./app"]
