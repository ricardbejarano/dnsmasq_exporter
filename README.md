<p align=center><img src=https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/160/apple/198/microscope_1f52c.png width=120px></p>
<h1 align=center>dnsmasq_exporter</h1>
<p align=center>A <a href=https://prometheus.io/>Prometheus</a>-style metrics <a href=https://prometheus.io/docs/instrumenting/exporters/>exporter</a> for the <a href=http://www.thekelleys.org.uk/dnsmasq/doc.html>Dnsmasq DNS forwarder</a></p>

***Note:** requires Dnsmasq version 2.69 or above!*


# Building

## Prerequisites

- The [Go](https://golang.org/doc/install) toolchain.
- The [miekg/dns](https://github.com/miekg/dns) library: `go get github.com/miekg/dns`
- The [prometheus/client_golang](https://github.com/prometheus/client_golang) library: `go get github.com/prometheus/client_golang`

## Building

```bash
go build dnsmasq_exporter.go
```

If successful, you'll find your `dnsmasq_exporter` binary on your working directory.

## Configuring

All configuration is done through environment variables:

- `DNSMASQ_SERVERS`: comma-separated list of the `<address>:<port>` pairs of the Dnsmasq servers
- `LISTEN_ADDR`: *(default: `0.0.0.0`)*, bind address for dnsmasq_exporter
- `LISTEN_PORT`: *(default: `9153`)*, bind port for dnsmasq_exporter

### Adding to Prometheus

Add the following to your Prometheus configuration file:

```yaml
scrape_configs:
  - job_name: dnsmasq
    static_configs:
      - targets: ['localhost:9153']
```

## Running

```bash
# monitor one server [10.0.0.1:53]
DNSMASQ_SERVERS=10.0.0.1:53 ./dnsmasq_exporter

# monitor many servers [10.0.0.1:53, 192.168.1.1:53]
DNSMASQ_SERVERS=10.0.0.1:53,192.168.1.1:53 ./dnsmasq_exporter
```


# Building (with Docker)

## Prerequisites

- A running [Docker engine](https://docs.docker.com/engine/)

## Building

- To build the [`glibc`](https://www.gnu.org/software/libc/)-based image: `docker build -t dnsmasq_exporter:glibc -f Dockerfile.glibc .`
- To build the [`musl`](https://www.musl-libc.org/)-based image: `docker build -t dnsmasq_exporter:musl -f Dockerfile.musl .`

### Official images

Instead of building your own, you can use the official dnsmasq_exporter Docker images.

#### Docker Hub

Available on [Docker Hub](https://hub.docker.com) as [`ricardbejarano/dnsmasq_exporter`](https://hub.docker.com/r/ricardbejarano/dnsmasq_exporter):

- [`1.0-glibc`, `1.0`, `glibc`, `master`, `latest` *(Dockerfile.glibc)*](https://github.com/ricardbejarano/dnsmasq_exporter/blob/master/Dockerfile.glibc)
- [`1.0-musl`, `musl` *(Dockerfile.musl)*](https://github.com/ricardbejarano/dnsmasq_exporter/blob/master/Dockerfile.musl)

#### Quay

Available on [Quay](https://quay.io) as:

- [`quay.io/1.0-glibc`](https://quay.io/repository/ricardbejarano/dnsmasq_exporter-glibc), tags: [`1.0`, `master`, `latest` *(Dockerfile.glibc)*](https://github.com/ricardbejarano/dnsmasq_exporter/blob/master/Dockerfile.glibc)
- [`quay.io/1.0-musl`](https://quay.io/repository/ricardbejarano/dnsmasq_exporter-musl), tags: [`1.0`, `master`, `latest` *(Dockerfile.musl)*](https://github.com/ricardbejarano/dnsmasq_exporter/blob/master/Dockerfile.musl)

## Running

```bash
# monitor one server [10.0.0.1:53]
docker run -it -p 9153:9153 -e 'DNSMASQ_SERVERS=10.0.0.1:53' ricardbejarano/dnsmasq_exporter

# monitor many servers [10.0.0.1:53, 192.168.1.1:53]
docker run -it -p 9153:9153 -e 'DNSMASQ_SERVERS=10.0.0.1:53,192.168.1.1:53' ricardbejarano/dnsmasq_exporter
```


# License

See [LICENSE](https://github.com/ricardbejarano/dnsmasq_exporter/blob/master/LICENSE).
