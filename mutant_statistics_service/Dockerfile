FROM golang:1.14-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /go/src/meli-exercise/mutant_statistics_service/

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod download

COPY . .

RUN go build -o compiled


FROM alpine:3.9

COPY --from=build_base /go/src/meli-exercise/mutant_statistics_service/compiled /dist/mutant_statistics_service

CMD ["/dist/mutant_statistics_service"]
