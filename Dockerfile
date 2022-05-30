FROM golang:1.18-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o main

# # # #
FROM golang:1.18-alpine
WORKDIR /app
COPY --from=builder /app .

EXPOSE 9000

CMD ./main