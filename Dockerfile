FROM golang:alpine as build

RUN apk update
RUN apk add git gcc
ADD . /app
WORKDIR /app
RUN go build -ldflags="-s -w" -v .

FROM scratch
COPY --from=build /app /app
EXPOSE 80
WORKDIR /app
ENTRYPOINT [ "/app/atomic-counter" ]
