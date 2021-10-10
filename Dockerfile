FROM golang:1.16 AS build

RUN apt-get update && apt-get install -y ca-certificates openssl

ARG cert_location=/usr/local/share/ca-certificates

RUN openssl s_client -showcerts -connect slack.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt

RUN update-ca-certificates

WORKDIR /src/

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o /bin/slackgo-api

FROM scratch

COPY --from=build /bin/slackgo-api  /bin/slackgo-api

ENTRYPOINT [ "/bin/slackgo-api" ]