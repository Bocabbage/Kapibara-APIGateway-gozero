FROM golang:1.21 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod tidy
COPY . .
COPY restful/auth/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/restful-auth restful/auth/auth.go


# FROM scratch
FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /app/restful-auth /app/restful-auth
COPY --from=builder /app/etc /app/etc

EXPOSE 443

# CMD ["./restful-auth", "-f", "etc/auth-api-localdemo.yaml"]
