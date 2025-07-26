ARG GO_VER
ARG ALPINE_VER
ARG GO_IMAGE=golang
ARG ALPINE_IMAGE=alpine

FROM ${GO_IMAGE}:${GO_VER}-alpine${ALPINE_VER} as builder
ARG VERSION
WORKDIR /src

RUN apk add git

# For caching
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app -ldflags "-X 'main.version=${VERSION}'" ./cmd/reports-api

FROM ${ALPINE_IMAGE}:${ALPINE_VER} AS app
WORKDIR /root/
COPY --from=builder /src/app .
CMD ["./app"]
