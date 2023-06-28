FROM golang:1.18 as builder

WORKDIR /app

RUN wget "https://github.com/go-swagger/go-swagger/releases/download/v0.30.3/swagger_linux_amd64" && \
    mv swagger_linux_amd64 /usr/local/bin/swagger && \
    chmod a+x /usr/local/bin/swagger

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go env -w GOPRIVATE=https://github.com/kurirmoo/*

COPY ./ /app

RUN wget "https://github.com/go-swagger/go-swagger/releases/download/v0.30.3/swagger_linux_amd64" && \
    mv swagger_linux_amd64 /usr/local/bin/swagger && \
    chmod a+x /usr/local/bin/swagger

RUN make all

FROM alpine:3.17.0

WORKDIR /app

COPY --from=builder /app/kurirmoo-server /app/kurirmoo-server

EXPOSE 8080

CMD [ "./kurirmoo-server", "--port=8080", "--host=0.0.0.0" ]