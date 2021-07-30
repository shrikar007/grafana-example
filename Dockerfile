FROM golang:alpine as build
COPY . /go/src/grafana-example
WORKDIR /go/src/grafana-example
RUN go mod vendor
RUN cd cmd && go build -o grafana-app .

FROM alpine:3.12
COPY --from=build /go/src/grafana-example/cmd/grafana-app .
EXPOSE 3000
CMD ["./grafana-app"]
