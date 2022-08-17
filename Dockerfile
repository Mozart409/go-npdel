FROM golang:1.18 as Builder

LABEL org.opencontainers.image.authors="amadeus@mozart409.com"
LABEL org.opencontainers.image.created="$(date -u +'%Y-%m-%dT%H:%M:%SZ')"

WORKDIR /app
COPY . .
RUN go mod tidy

RUN go build -o /app/go-npdel

FROM scratch as Runner
WORKDIR /app
COPY --from=Builder /app/go-npdel .
ENTRYPOINT ["/app/go-npdel"]