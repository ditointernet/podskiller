FROM golang:1.13-alpine3.10 as builder

ENV CGO_ENABLED=0

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o /podskiller -mod=readonly ./main.go

FROM alpine

COPY --from=builder /podskiller /podskiller

ENTRYPOINT [ "/podskiller" ]