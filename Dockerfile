FROM golang:1.22 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

RUN mkdir /config

COPY . .
RUN CGO_ENABLED=0 go build -o dieppe

FROM cgr.dev/chainguard/static

COPY --from=builder /build/dieppe /dieppe
COPY --from=builder /config /config

EXPOSE 8000

ENTRYPOINT ["/dieppe"]
CMD ["start", "--config=/config"]
