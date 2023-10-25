#  deploy-builder
# --------------------------------------------------

FROM golang:1.21.1 as deploy-builder

ENV CGO_ENABLED 0

ARG VERSION

WORKDIR /workspace

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg \
    go mod download

COPY app/ ./app
RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags "-s -w -X main.Version=${VERSION}" -o app ./app

#  deploy
# --------------------------------------------------

FROM debian:bookworm-20230320-slim as deploy

RUN apt-get update \
    && rm -rf /var/lib/apt/lists/*

COPY --from=deploy-builder /workspace/app .

CMD [ "./app" ]

#  dev
# --------------------------------------------------

FROM golang:1.21.1 as dev

WORKDIR /workspace

RUN go install github.com/go-delve/delve/cmd/dlv@v1.21.1
RUN go install github.com/cosmtrek/air@v1.45.0

CMD [ "air", "-c", ".air.toml" ]
