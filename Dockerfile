ARG version=alpine
FROM golang:${version} AS builder

WORKDIR /app
COPY app.go .
RUN go build app.go


FROM scratch
LABEL org.opencontainers.image.authors="tin3ga" \
      org.opencontainers.image.description="Container image for https://github.com/tin3ga/testEnv"

COPY --from=builder /app/app .
ENV NAME="Jane Doe"
ENV CHARACTER="SpiderMan"
ENV RUN_TIME="0"

CMD ["./app"]