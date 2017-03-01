FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/yujie0121/canvas

ENV USER root
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET jc4lKaPPMBh7bxAb

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://root@localhost:5432/canvas?sslmode=disable

WORKDIR /go/src/github.com/yujie0121/canvas

RUN godep go build

EXPOSE 8888
CMD ./canvas