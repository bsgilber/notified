# Stage 1
FROM golang:alpine3.14  AS build
WORKDIR /app/src/
RUN apk add --update make
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make test build

# Stage 2
FROM alpine:3.14 AS run
WORKDIR /app
COPY --from=build /app/src/dist/notified .

EXPOSE 7077
ENTRYPOINT [ "/app/notified" ]
