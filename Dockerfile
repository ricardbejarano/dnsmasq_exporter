FROM golang:1-alpine AS build

COPY . /build
RUN apk add make && \
    cd /build && \
      make

RUN mkdir -p /rootfs/bin && \
      cp /build/bin/dnsmasq_exporter /rootfs/bin/ && \
    mkdir -p /rootfs/etc && \
      echo "nogroup:*:10000:nobody" > /rootfs/etc/group && \
      echo "nobody:*:10000:10000:::" > /rootfs/etc/passwd


FROM scratch

COPY --from=build --chown=10000:10000 /rootfs /

ENV EXPORTER_LISTEN_ADDR="0.0.0.0:9153"
USER 10000:10000
EXPOSE 9153/tcp
ENTRYPOINT ["/bin/dnsmasq_exporter"]
