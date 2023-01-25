# build the manager from a dev image
FROM --platform=$BUILDPLATFORM golang:alpine as build

ARG TARGETOS
ARG TARGETARCH

RUN apk --no-cache add tzdata

WORKDIR /theway
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -installsuffix cgo -o theway .

FROM gcr.io/distroless/static AS final

USER nonroot:nonroot

# copy compiled app
COPY --from=build --chown=nonroot:nonroot /theway/theway /theway

# run binary
ENTRYPOINT ["/theway"]
