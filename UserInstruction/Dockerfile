FROM golang:1.18-alpine

RUN mkdir /app

RUN addgroup -S pantigroup
RUN adduser -S -D -h /app pantiuser pantigroup
RUN chown -R pantiuser:pantigroup /app
USER pantiuser

COPY main.go /app

EXPOSE 8080
CMD go run /app/main.go