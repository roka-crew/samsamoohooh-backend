# 빌드 스테이지
FROM golang:1.24-alpine AS builder

# 필요한 빌드 도구 설치
RUN apk add --no-cache git

# 작업 디렉토리 설정
WORKDIR /app

# 의존성 파일 복사 및 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드 복사
COPY . .

# 애플리케이션 빌드
RUN go build -o bin/samsamoohooh ./cmd/samsamoohooh/samsamoohooh.go

# 실행 스테이지
FROM alpine:3.19

# 실행에 필요한 timezone 데이터 설치
RUN apk add --no-cache tzdata

# 작업 디렉토리 설정
WORKDIR /app

# 빌드 스테이지에서 빌드된 바이너리 복사
COPY --from=builder /app/bin/samsamoohooh /app/samsamoohooh

# 설정 파일 복사
COPY configs/config.yaml /app/configs/config.yaml

# 실행 권한 설정
RUN chmod +x /app/samsamoohooh

# 컨테이너가 리스닝할 포트 노출
EXPOSE 80

# 애플리케이션 실행
CMD ["/app/samsamoohooh"]