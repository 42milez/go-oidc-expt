#  deploy-builder
# --------------------------------------------------

FROM golang:1.22.3 as deploy-builder

ENV CGO_ENABLED 0

ARG VERSION

WORKDIR /workspace

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg \
    go mod download

COPY app/idp/ ./app/idp
COPY app/pkg/ ./app/pkg

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags "-s -w -X main.version=${VERSION}" -o idp ./app/idp

#  deploy
# --------------------------------------------------

FROM debian:bookworm-20230320-slim as deploy

RUN apt-get update \
    && rm -rf /var/lib/apt/lists/*

COPY --from=deploy-builder /workspace/idp .

CMD [ "./idp" ]

#  dev
# --------------------------------------------------

FROM golang:1.22.3 as dev

WORKDIR /workspace

RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.1
RUN go install github.com/cosmtrek/air@v1.45.0

CMD [ "air", "-c", ".air.toml" ]
