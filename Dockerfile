FROM golang:alpine as build

RUN apk update
RUN apk add git
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -a -ldflags="-s -w" -v -o main .

FROM scratch
WORKDIR /usr/src/app
COPY --from=build /app .
EXPOSE 80
ENTRYPOINT ["/usr/src/app/main"]
