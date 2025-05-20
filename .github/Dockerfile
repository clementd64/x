FROM stagex/pallet-go AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM deps AS build
ARG APP
COPY . .
RUN --network=none <<-EOF
    set -eu
    go build -trimpath -ldflags "${LDFLAGS}" -o app "./${APP}"
    install -Dm755 -t /rootfs/usr/bin app
EOF

FROM stagex/core-filesystem
COPY --from=stagex/core-ca-certificates . /
COPY --from=build /rootfs/ /
ENTRYPOINT [ "app" ]
