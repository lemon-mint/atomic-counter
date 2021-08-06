FROM golang:alpine as build

RUN apk update
RUN apk add git
ADD . /app
WORKDIR /app
RUN go build -ldflags="-s -w" -v .

FROM alpine:latest
COPY --from=build /app /app
EXPOSE 80
WORKDIR /app
CMD /app/atomic-counter
