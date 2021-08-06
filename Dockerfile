FROM golang:alpine as build

RUN apk update
RUN apk add git gcc
ADD . /app
WORKDIR /app
RUN go build -ldflags="-s -w" -v -o server .

FROM scratch
COPY --from=build /app/server /server
EXPOSE 80
ENTRYPOINT [ "server" ]
