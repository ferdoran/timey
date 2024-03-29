FROM golang:1.18-bullseye as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...
RUN go vet -v
RUN go test -v ./...

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static
COPY --from=build /go/bin/app /
CMD ["/app"]