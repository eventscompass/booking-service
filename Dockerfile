# Build
FROM golang:1.21.2-alpine as builder

ENV GO111MODULE=on
ENV GOPRIVATE=github.com/eventscompass
ENV GOOS=linux

WORKDIR /service
COPY . .

RUN CGO_ENABLED=0 go build -o "/tmp/bookingservice" ./src

#####

# Run
FROM scratch
COPY --from=builder /tmp/bookingservice .
EXPOSE 8080
CMD [ "./bookingservice" ]