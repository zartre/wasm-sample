FROM golang:1.17.3-alpine AS build-stage

WORKDIR /app

COPY src/go ./
RUN cd main && GOOS=js GOARCH=wasm go build -o ../out/main.wasm

WORKDIR /app
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" out/

FROM node:alpine

WORKDIR /app

COPY --from=build-stage /app/out ./
COPY src/web ./

EXPOSE 8080

CMD npx http-server