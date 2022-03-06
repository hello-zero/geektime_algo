FROM golang:1.17 AS build
RUN mkdir /httpserver
WORKDIR /httpserver/
COPY . .
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o ./httpserver ./main.go

FROM busybox
COPY --from=build /httpserver/httpserver /httpserver/httpserver
WORKDIR /httpserver
EXPOSE 80
ENTRYPOINT ["./httpserver"]