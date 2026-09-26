FROM golang:1.27-alpine AS build
WORKDIR /src
COPY apps/api/go.mod apps/api/go.sum ./
RUN go mod download
COPY apps/api/ ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api

FROM golang:1.27-alpine AS goose-build
RUN CGO_ENABLED=0 go install -tags='no_clickhouse no_libsql no_mssql no_mysql no_sqlite3 no_vertica no_ydb' github.com/pressly/goose/v3/cmd/goose@v3.28.0

FROM alpine:3.23 AS migrate
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=goose-build /go/bin/goose /usr/local/bin/goose
COPY apps/api/db/migrations ./db/migrations
USER 65532:65532
ENTRYPOINT ["goose", "-dir", "/app/db/migrations"]

FROM alpine:3.23 AS api
RUN apk add --no-cache ca-certificates
COPY --from=build /out/api /usr/local/bin/api
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["api"]
