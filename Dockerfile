FROM golang:1.17-alpine as builder

RUN mkdir /app   
COPY . /app/
WORKDIR /app 

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o /app/main .

FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY --from=builder /app/main /
EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/main"]

