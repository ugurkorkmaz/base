FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY ./backend /app/backend 

WORKDIR /app

RUN ls -la

EXPOSE 8080

ENTRYPOINT [ "./backend" ]