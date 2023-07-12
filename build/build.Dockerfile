FROM golang:1.20 as base

FROM base as built

WORKDIR /go/app/wc3_game_tracker
COPY .. .

ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./wc3_game_tracker ./*.go


FROM busybox

WORKDIR ./usr/bin/

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=built /go/app/wc3_game_tracker/wc3_game_tracker .
COPY --from=built /go/app/wc3_game_tracker/templates ./templates

CMD ["./wc3_game_tracker"]