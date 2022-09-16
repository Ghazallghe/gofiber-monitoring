FROM golang:latest AS builder

WORKDIR /app

ENV GOPROXY=https://goproxy.io,direct 

# make a change to binaray, so it can be run on alpine container
ENV CGO_ENABLED=0 

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY main.go ./
COPY pkg ./pkg

RUN go build -o monitor -tags netgo -a -v



FROM alpine:3.16.2 AS runner
WORKDIR /app

RUN adduser --system --uid 1001 application
RUN addgroup --system --gid 1001 application
USER application

WORKDIR /app

COPY --from=builder --chown=application:application /app/monitor ./
COPY --chown=application:application .env ./

CMD ["./monitor", "serve"]
