FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY ./todo /app/todo 

WORKDIR /app

RUN ls -la

EXPOSE 8080

ENTRYPOINT [ "./todo" ]