<p align="center"><img src="https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/160/apple/198/microscope_1f52c.png" width="120px"></p>
<h1 align="center">dnsmasq_exporter</h1>
<p align="center">A <a href="https://prometheus.io/">Prometheus</a>-style metrics <a href="https://prometheus.io/docs/instrumenting/exporters/">exporter</a> for the <a href="http://www.thekelleys.org.uk/dnsmasq/doc.html">Dnsmasq DNS forwarder</a></p>


Requires **Dnsmasq 2.69 or above**!


## Tags

### Docker Hub

Available on [Docker Hub](https://hub.docker.com) as [`ricardbejarano/dnsmasq_exporter`](https://hub.docker.com/r/ricardbejarano/dnsmasq_exporter):

- [`1.2-glibc`, `1.2`, `glibc`, `master`, `latest` *(Dockerfile.glibc)*](https://github.com/ricardbejarano/dnsmasq_exporter/blob/master/Dockerfile.glibc)
- [`1.2-musl`, `musl` *(Dockerfile.musl)*](https://github.com/ricardbejarano/dnsmasq_exporter/blob/master/Dockerfile.musl)

### Quay

Available on [Quay](https://quay.io) as:

- [`quay.io/ricardbejarano/dnsmasq_exporter-glibc`](https://quay.io/repository/ricardbejarano/dnsmasq_exporter-glibc), tags: [`1.2`, `master`, `latest` *(Dockerfile.glibc)*](https://github.com/ricardbejarano/dnsmasq_exporter/blob/master/Dockerfile.glibc)
- [`quay.io/ricardbejarano/dnsmasq_exporter-musl`](https://quay.io/repository/ricardbejarano/dnsmasq_exporter-musl), tags: [`1.2`, `master`, `latest` *(Dockerfile.musl)*](https://github.com/ricardbejarano/dnsmasq_exporter/blob/master/Dockerfile.musl)


## Features

* Super tiny (about `13.2MB`)
* Compiled from source during build time
* Built `FROM scratch`, with zero bloat (see [Filesystem](#filesystem))
* Reduced attack surface (no shell, no UNIX tools, no package manager...)
* Runs as unprivileged (non-`root`) user


## Building

### Local

```bash
cd dnsmasq_exporter
go get \
  github.com/miekg/dns \
  github.com/prometheus/client_golang/prometheus
go build
```

### Docker

- To build the `glibc`-based image: `$ docker build -t dnsmasq_exporter:glibc -f Dockerfile.glibc .`
- To build the `musl`-based image: `$ docker build -t dnsmasq_exporter:musl -f Dockerfile.musl .`


## Configuration

### Environment variables

- `DNSMASQ_SERVERS`: *(default: `127.0.0.1:53`)* comma-separated list of `<address>:<port>` pairs of Dnsmasq servers' addresses
- `LISTEN_ADDR`: *(default: `0.0.0.0`)*, bind address for `dnsmasq_exporter`
- `LISTEN_PORT`: *(default: `9153`)*, bind port for `dnsmasq_exporter`

### Adding to Prometheus

Add the following to `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: dnsmasq
    static_configs:
      - targets: ['<DNSMASQ_EXPORTER_ADDRESS>:9153']
```


## Filesystem

```
/
├── dnsmasq_exporter
└── etc/
    ├── group
    └── passwd
```


## License

See [LICENSE](https://github.com/ricardbejarano/dnsmasq_exporter/blob/master/LICENSE).
