FROM golang:1.14-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /go/src/meli-exercise/mutant_service/

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

COPY . .

RUN go mod download

RUN go build -o compiled


FROM alpine:3.9

COPY --from=build_base /go/src/meli-exercise/mutant_service/compiled /dist/mutant_service

CMD ["/dist/mutant_service"]
