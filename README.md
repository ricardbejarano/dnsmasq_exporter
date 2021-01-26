<p align="center"><img src="https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/160/apple/198/microscope_1f52c.png" width="120px"></p>
<h1 align="center">dnsmasq_exporter</h1>
<p align="center"><a href="https://prometheus.io/">Prometheus</a> metrics <a href="https://prometheus.io/docs/instrumenting/exporters/">exporter</a> for the <a href="http://www.thekelleys.org.uk/dnsmasq/doc.html">Dnsmasq</a> network services</p>


# Description

Prometheus metrics exporter for the Dnsmasq network services.

Metrics are obtained by querying Dnsmasq directly over its DNS endpoint.


# Usage

## Considerations

* Requires **Dnsmasq 2.69 or above**!
* **Tune your configuration**. All configuration is done through *environment variables*:
  * `DNSMASQ_SERVERS`: *(defaults to `127.0.0.1:53`)* comma-separated list of `<address>:<port>` pairs of Dnsmasq servers' addresses
  * `EXPORTER_LISTEN_ADDR`: *(defaults to `127.0.0.1:9153`)*, bind address for `dnsmasq_exporter`

## With the prebuilt container image

Available on [Docker Hub](https://hub.docker.com) as [`docker.io/ricardbejarano/dnsmasq_exporter`](https://hub.docker.com/r/ricardbejarano/dnsmasq_exporter):

- [`1.4`, `latest` *(Dockerfile)*](Dockerfile)

Also available on [Quay](https://quay.io) as [`quay.io/ricardbejarano/dnsmasq_exporter`](https://quay.io/repository/ricardbejarano/dnsmasq_exporter):

- [`1.4`, `latest` *(Dockerfile)*](Dockerfile)

Any of both registries will do, example:

```bash
docker run -it -p 9153:9153 quay.io/ricardbejarano/dnsmasq_exporter
```

## Building the container image from source

First clone the repository, and `cd` into it:

```bash
docker build -t dnsmasq_exporter .
```

Now run it:

```bash
docker run -it -p 9153:9153 dnsmasq_exporter
```

## Building the binary from source

First clone the repository, and `cd` into it.

```bash
make
```

Now run it:

```bash
./bin/dnsmasq_exporter
```

## Integrating with Prometheus

Add the following to `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: "dnsmasq"
    static_configs:
      - targets: ["<DNSMASQ_EXPORTER_ADDRESS>:9153"]
```


# License

MIT licensed, see [LICENSE](LICENSE) for more details.
