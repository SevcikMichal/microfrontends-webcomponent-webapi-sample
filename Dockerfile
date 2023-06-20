FROM golang:1.19.4 AS build
WORKDIR /app
COPY models ./models
COPY rest-api ./rest-api
COPY router ./router
COPY db-service ./db-service
COPY main.go .

ENV CGO_ENABLED=0

RUN go mod init ambulance-webapi
RUN go mod tidy

RUN go build -a -o ambulance-webapi .

FROM scratch
COPY --from=build /app/ambulance-webapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./ambulance-webapi"]
