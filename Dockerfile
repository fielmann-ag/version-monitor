FROM golang as builder
WORKDIR /src/

ARG COMMAND

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o bin/app main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

ARG COMMAND

RUN addgroup -S "app" && adduser -S -D -u 1000 --gecos "" "app" -G "app"
USER app
WORKDIR /home/app
ENV COMMAND=app

COPY --from=builder --chown=app:app /src/bin/app /home/app/app

CMD ["/home/app/app"]
