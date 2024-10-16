FROM cimg/go:1.23 AS builder

COPY --chown=circleci:circleci ./ .

RUN go build -o ./app ./main.go
RUN curl https://upload.wikimedia.org/wikipedia/commons/3/3f/JPEG_example_flower.jpg -o ./file/test.jpg

FROM debian:trixie-slim
COPY --from=builder /home/circleci/project .

EXPOSE 8080
ENTRYPOINT ["./app"]

