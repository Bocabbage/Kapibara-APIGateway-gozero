FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY restful/auth/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/restful-auth restful/auth/auth.go


# FROM scratch
FROM ubuntu

WORKDIR /app
COPY --from=builder /app/restful-auth /app/restful-auth
COPY --from=builder /app/etc /app/etc

EXPOSE 443

# CMD ["./restful-auth", "-f", "etc/auth-api-localdemo.yaml"]
