# 빌드 스테이지
FROM golang:1.24 AS builder

# CA 인증서 설치
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

# 작업 디렉토리 설정
WORKDIR /app

# 로컬 모듈 복사
COPY ../ferrari-common /go/src/git.datau.co.kr/ferrari/ferrari-common
COPY ../benz-common /go/src/git.datau.co.kr/benz/benz-common
COPY ../earth-common /go/src/git.datau.co.kr/earth/earth-common

# go.mod와 go.sum 복사 및 의존성 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드 복사
COPY . .

# Go 애플리케이션 빌드
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -o myapp main.go

# 런타임 스테이지
FROM alpine:latest
RUN apk add --no-cache tzdata ca-certificates
ENV TZ=Asia/Seoul
WORKDIR /root/
COPY --from=builder /app/myapp .
ENV BENTLEY=false
ENV BENZ=true
ENV FERRARI=false
ENV TESLA=false
ENV MARS=false
ENV SATURN=false
CMD ["./myapp"]