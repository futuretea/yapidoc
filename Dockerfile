FROM golang:1.11 AS builder
ADD ./ /src
WORKDIR /src
RUN make

FROM alpine:3.9.2
RUN sed -i "s|http://dl-cdn.alpinelinux.org|https://mirrors.aliyun.com|g" /etc/apk/repositories \
    && apk update \
    && apk add tzdata
ENV TZ=Asia/Shanghai
RUN mkdir -p /app
WORKDIR /app
COPY templates ./templates
COPY config.yaml ./config.yaml
COPY --from=builder /src/yapidoc .
CMD [ "./yapidoc", "--config", "./config.yaml"]

