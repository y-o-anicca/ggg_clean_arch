FROM golang:1.18.3-bullseye as build
ENV TZ=Asia/Tokyo

WORKDIR /go/src/app
COPY . .

WORKDIR /go/src/app
RUN go build -o /go/bin/app ./cmd/api/main.go

FROM gcr.io/distroless/base-debian11
ENV TZ=Asia/Tokyo
COPY --from=build /go/bin/app /
CMD ["/app"]
