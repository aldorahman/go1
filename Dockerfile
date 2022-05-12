FROM golang:1.17-alpine as builder
WORKDIR /belajar-API
COPY . .
RUN go build -o main

# FROM alpine:3.13
# WORKDIR /belajar-API
# COPY --from=builder /belajar-API .

EXPOSE 9000

CMD ./main