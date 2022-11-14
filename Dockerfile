FROM golang:1.18.8-alpine3.16 AS server
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o onlien-json-diff cmd/main.go

FROM node:16.18-alpine AS client
WORKDIR /app
COPY web/package.json .
COPY web/yarn.lock .
RUN yarn install
COPY web .
RUN yarn run build

FROM alpine:3.16.3
WORKDIR /app
COPY --from=server /app/onlien-json-diff ./
COPY --from=client /app/build ./web
CMD ["sh", "-c", "/app/onlien-json-diff -server.env=prod"]
