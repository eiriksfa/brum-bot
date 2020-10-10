FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main ./cmd/brum-bot
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./main"]