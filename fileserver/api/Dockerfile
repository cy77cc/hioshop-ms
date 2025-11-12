FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod tidy
COPY . .

RUN go build -ldflags="-s -w" -o /app/file app/fileserver/api/file.go


FROM alpine
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/file /app/file
COPY app/fileserver/api/etc /app/etc

EXPOSE 39000

CMD ["./file", "-f", "etc/file.yaml"]
