FROM golang:1.20-bullseye as builder
WORKDIR /go/simple_api
COPY main.go .

RUN GOOS=linux GOARCH=amd64 go build main.go


FROM debian:bullseye-slim
WORKDIR /simple_api
COPY --from=builder /go/simple_api .
COPY ./tools/entrypoint.sh .
RUN chmod 755 ./entrypoint.sh

EXPOSE 8081
ENTRYPOINT ["./entrypoint.sh"]
