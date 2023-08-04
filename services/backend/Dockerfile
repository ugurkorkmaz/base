FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY ./backend /app/backend 

WORKDIR /app

EXPOSE 8080

ENTRYPOINT [ "./backend" ]