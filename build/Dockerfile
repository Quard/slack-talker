FROM golang:1.12-alpine as build
WORKDIR /app
COPY . .
RUN apk add git
ENV GO111MODULE on
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -v -o /slack-talker cmd/slack-talker/main.go

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=build /slack-talker /slack-talker
CMD /slack-talker
