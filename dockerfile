FROM golang:1.20-alpine3.17 as builder

RUN mkdir /api
WORKDIR /api
ADD . /api

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.2
RUN swag init
RUN go test ./...
RUN go build -o /app/api .

FROM alpine:3.17

COPY --from=builder /app/api /api
COPY --from=builder /api/docs /docs
ADD /configs /configs
ADD /assets /assets

ENV TZ=Asia/Bangkok

EXPOSE 8000

ENTRYPOINT ["/api"]
