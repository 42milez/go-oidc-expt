#  deploy-builder
# --------------------------------------------------

FROM golang:1.20.2 as deploy-builder

ENV CGO_ENABLED 0

ARG VERSION

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY idp ./
RUN go build -trimpath -ldflags "-s -w -X main.version=${VERSION}" -o app ./app

#  deploy
# --------------------------------------------------

FROM debian:bookworm-20230320-slim as deploy

RUN apt-get update \
    && rm -rf /var/lib/apt/lists/*

COPY --from=deploy-builder /workspace/app .

CMD [ "./app" ]

#  dev
# --------------------------------------------------

FROM golang:1.20.2 as dev

WORKDIR /workspace

RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.2
RUN go install github.com/cosmtrek/air@v1.44.0

CMD [ "air", "-c", ".air.toml" ]
