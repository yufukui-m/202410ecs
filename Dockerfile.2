FROM cimg/go:1.23 AS builder

COPY ./ .
RUN go build -o ./app ./main.go

FROM debian:trixie-slim
COPY --from=builder /home/circleci/project .

EXPOSE 8080
ENTRYPOINT ["./app"]

