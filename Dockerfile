FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /gotter

WORKDIR /gotter

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.7


ENV PG_HOST=""
ENV PG_PORT=""
ENV PG_USER=""
ENV PG_DB=""
ENV PG_PASSWORD=""

ENV JWT_SECRET=""

ENV PORT=""

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /gotter/api_server /usr/bin/api_server

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]