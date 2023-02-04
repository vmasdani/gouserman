FROM golang:1.20.0-alpine3.16 as backend-env
WORKDIR /app

COPY ./backend/. .
RUN --mount=type=cache,target=/go/pkg/mod go get
RUN --mount=type=cache,target=/go/pkg/mod go build -o gouserman

FROM node:18.14.0-alpine3.17 as frontend-env
WORKDIR /app
COPY ./frontend/. .
RUN --mount=type=cache,target=/root/.yarn YARN_CACHE_FOLDER=/root/.yarn yarn install
RUN --mount=type=cache,target=/root/.yarn YARN_CACHE_FOLDER=/root/.yarn yarn build

FROM alpine:3.17.1
WORKDIR /app
COPY --from=backend-env /app/ .
COPY --from=frontend-env /app/dist/. frontend/
CMD [ "./gouserman" ]