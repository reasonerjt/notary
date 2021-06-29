ARG VERSION=1.15.6
FROM golang:${VERSION}-alpine
ARG VERSION=1.15.6
ENV NOTARYPKG github.com/theupdateframework/notary
COPY . /go/src/${NOTARYPKG}
WORKDIR /go/src/${NOTARYPKG}
RUN go install -ldflags "-X main.goversion=$VERSION" ${NOTARYPKG}/cmd/poc

ENTRYPOINT [ "poc" ]