FROM docker.io/golang:1 AS build

RUN mkdir -p /rootfs/etc

RUN echo "nobody:*:10000:nobody" > /rootfs/etc/group \
    && echo "nobody:*:10000:10000:::" > /rootfs/etc/passwd

COPY . /build
RUN cd /build \
    && CGO_ENABLED=0 go build -o bin/ . \
    && mkdir -p /rootfs \
    && cp -r /build/bin /rootfs/


FROM scratch

COPY --from=build --chown=10000:10000 /rootfs /

ENV EXPORTER_LISTEN_ADDR="0.0.0.0:9153"
USER nobody:nobody
WORKDIR /
EXPOSE 9153/TCP
ENTRYPOINT ["/bin/dnsmasq_exporter"]
