# FROM: 베이스 이미지 생성
FROM golang:1.20-alpine

# RUN: `이미지를 생성하는 과정`에서 사용할 명령문 실행
RUN apk add --no-cache bash

# WORKDIT: 작업 디렉토리를 지정
WORKDIR /go-dir

# COPY: 파일 복사 (이동)
COPY ./ ./

# ENTRYPOINT: `컨테이너가 시작`할 때 실행되는 명령어
ENTRYPOINT ["/bin/bash", "-c", "sleep 500"]

# EXPOSE: 컨테이너 내부에서 사용 중인 포트를 문서화하기