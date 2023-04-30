FROM golang:1.20 AS Builder

WORKDIR /citest
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /cmd .

FROM alpine:latest
COPY --from=builder /cmd /cmd
CMD ["./cmd"]