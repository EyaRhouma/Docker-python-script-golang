FROM golang:alpine AS build

RUN apk update && apk add make gcc libc-dev

WORKDIR /go/src/solver
ADD test.go .
RUN go build -o main test.go

FROM alpine
WORKDIR /app
COPY --from=build /go/src/solver/main /app/
ENTRYPOINT ["/app/main"]
