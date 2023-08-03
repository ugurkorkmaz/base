# Builder stage
FROM node:latest as builder
WORKDIR /app
COPY package*.json .
RUN npm install
COPY . .
RUN npm run build
RUN npm prune --production


# Production stage
FROM gcr.io/distroless/nodejs20-debian11
WORKDIR /app

COPY --from=builder /app/build build/
COPY --from=builder /app/node_modules node_modules/
COPY package.json .

USER nonroot

EXPOSE 3000
CMD ["build"]