FROM golang:1.20 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

RUN mkdir /config

COPY . .
RUN CGO_ENABLED=0 go build -o dieppe

FROM scratch

COPY --from=builder /build/dieppe /dieppe
# Need to copy the config directory as the scratch image doesn't have mkdir
COPY --from=builder /config /config

EXPOSE 80

ENTRYPOINT ["/dieppe"]
CMD ["start", "--config=/config"]
