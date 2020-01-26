FROM golang:alpine as builder
RUN mkdir /swvl-notification-center 
ADD . /swvl-notification-center/ 
WORKDIR /swvl-notification-center 
RUN go build -o main . 
CMD ["/swvl-notification-center/main"]

FROM alpine
COPY --from=builder /swvl-notification-center/main /app/
WORKDIR /app
CMD ["./main"]

