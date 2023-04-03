FROM public.ecr.aws/docker/library/golang:1.20-alpine3.17 as builder

WORKDIR /app

COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/gk/ pkg/gk/
COPY vendor/ vendor/

COPY go.mod go.sum ./

# Just for debug purposes
#ARG TARGETPLATFORM
#RUN echo "$TARGETPLATFORM"
#ARG TARGETOS
#RUN echo "$TARGETOS"
#ARG TARGETARCH
#RUN echo "$TARGETARCH"
#ARG BUILDPLATFORM
#RUN echo "$BUILDPLATFORM"
#ARG BUILDOS
#RUN echo "$BUILDOS"
#ARG BUILDARCH
#RUN echo "$BUILDARCH"

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -o gatekeeper -ldflags="-s -w" ./cmd/main.go

FROM public.ecr.aws/docker/library/alpine:3.17

COPY --from=builder /app/gatekeeper /gatekeeper

CMD ["/gatekeeper"]