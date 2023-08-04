FROM alpine:latest as runtime

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY backend /app/backend

EXPOSE 8080
ENTRYPOINT [ "./backend" ]