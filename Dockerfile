FROM golang:1.19 as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o golangToAwsFromNone .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/golangToAwsFromNone .
EXPOSE 8080
CMD ["./golangToAwsFromNone"]
