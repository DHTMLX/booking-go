FROM debian:bookworm-slim

WORKDIR /app

COPY ./booking-go /app/booking-go

CMD ["/app/booking-go"]