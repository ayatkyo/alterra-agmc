# Builder Stage
FROM golang:1.19.1-alpine as Builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod download && go mod tidy
RUN go build -o agmc-ayat .

# App Stage
FROM alpine:latest
RUN apk update
WORKDIR /app
COPY --from=builder /app/agmc-ayat .
EXPOSE 15000
CMD [ "./agmc-ayat" ]