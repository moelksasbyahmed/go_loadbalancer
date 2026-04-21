FROM golang:1.26.2-alpine as builder 

WORKDIR /app


RUN apk add --no-cache make

COPY go.mod go.sum ./
RUN go mod download
COPY . . 

RUN make build

FROM alpine:latest 
WORKDIR /app
COPY --from=builder /app/Go_LoadBalancer .
COPY --from=builder /app/config.yaml .
COPY --from=builder /app/servers.yaml .
ENV PORT_ONE=5000
CMD  ./Go_LoadBalancer start -c ./config.yaml -p $PORT_ONE 
EXPOSE $PORT_ONE
EXPOSE 8086

