FROM golang:1.12 AS builder
WORKDIR /opt
COPY main.go /opt
COPY Makefile /opt
RUN make all

FROM debian:10
COPY --from=builder /opt/main /opt
WORKDIR /opt
EXPOSE 5000
CMD ["./main"]