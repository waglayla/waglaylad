ARG PYIPAD_IMAGE
ARG PYIPADMINER_IMAGE

FROM ${PYIPAD_IMAGE} as waglaylad
FROM ${PYIPADMINER_IMAGE} as waglaylaminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/waglayla/waglaylad

WORKDIR /go/src/github.com/waglayla/waglaylad

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=waglaylad /app/ /app/
COPY --from=waglaylaminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/waglayla/waglaylad/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]
