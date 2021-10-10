FROM golang:1.16 AS build

WORKDIR /src/

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags '-extldflags "-static"' -o /bin/slackgo-api

FROM scratch
ADD ssl/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/slackgo-api  /bin/slackgo-api

ENTRYPOINT [ "/bin/slackgo-api" ]