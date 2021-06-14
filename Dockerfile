FROM golang:1.16-alpine3.13 as builder
ENV GOPROXY=https://goproxy.io
WORKDIR /build
ADD . /build/
RUN CGO_ENABLED=0 go build -a -ldflags "-s -w" -o showip /build/

FROM busybox:1.32.0
COPY --from=builder /build/showip /
CMD ["/showip"]