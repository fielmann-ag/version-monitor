FROM golang as builder
WORKDIR /src/

ARG COMMAND

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o bin/version-monitor main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

ARG COMMAND

RUN addgroup -S "version-monitor" && adduser -S -D -u 1000 --gecos "" "version-monitor" -G "version-monitor"
USER version-monitor
WORKDIR /home/version-monitor
ENV COMMAND=version-monitor

COPY --from=builder --chown=version-monitor:version-monitor /src/bin/version-monitor /home/version-monitor/version-monitor

CMD ["/home/version-monitor/version-monitor"]
