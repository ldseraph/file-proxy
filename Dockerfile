FROM golang:1.19.6 AS BUILD
WORKDIR /go/src/file-proxy
COPY . .
RUN ./build.sh

FROM alpine:latest
WORKDIR /app
COPY --from=BUILD /go/src/file-proxy/file-proxy ./file-proxy
COPY ./file-proxy .
CMD ["file-proxy"]
EXPOSE 3000
