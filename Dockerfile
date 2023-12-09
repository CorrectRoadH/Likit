# Build frontend dist.
FROM node:18-alpine AS frontend
WORKDIR /frontend-build

COPY . .

WORKDIR /frontend-build/dashboard

RUN corepack enable && yarn

RUN yarn build

# Build backend exec file.
FROM golang:1.21-alpine AS backend
WORKDIR /backend-build

COPY . .
COPY --from=frontend /frontend-build/dashboard/dist ./internal/adapter/in/restful/dist

RUN CGO_ENABLED=0 go build -o likit ./main.go

# Make workspace with above generated files.
FROM alpine:latest AS monolithic
WORKDIR /usr/local/likit

RUN apk add --no-cache tzdata
ENV TZ="UTC"

COPY --from=backend /backend-build/likit /usr/local/likit/

EXPOSE 7789

# Directory to store the data, which can be referenced as the mounting point.
RUN mkdir -p /var/opt/likit
VOLUME /var/opt/likit

ENTRYPOINT ["./likit"]
