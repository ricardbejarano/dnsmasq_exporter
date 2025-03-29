<div align="center">
  <p><img src="https://em-content.zobj.net/thumbs/160/apple/391/fire_1f525.png" width="100px"></p>
  <h1>dnsmasq_exporter</h1>
  <p>Prometheus exporter for <a href="https://thekelleys.org.uk/dnsmasq/doc.html">Dnsmasq</a></p>
  <code>docker pull quay.io/ricardbejarano/dnsmasq_exporter</code>
</div>


## Usage

### Considerations

* **Requires Dnsmasq 2.69 or above!**
* **Tune your `dnsmasq_exporter` configuration**. All configuration is done through environment variables:
  * `DNSMASQ_SERVERS`: *(defaults to `127.0.0.1:53`)* comma-separated list of `<address>:<port>` pairs of Dnsmasq servers' addresses
  * `EXPORTER_LISTEN_ADDR`: *(defaults to `127.0.0.1:9153`)*, bind address for `dnsmasq_exporter`

### Using the official container image

#### Docker Hub

Available on Docker Hub as [`docker.io/ricardbejarano/dnsmasq_exporter`](https://hub.docker.com/r/ricardbejarano/dnsmasq_exporter):

- [`1.4.6`, `latest` *(Dockerfile)*](Dockerfile)

#### RedHat Quay

Available on RedHat Quay as [`quay.io/ricardbejarano/dnsmasq_exporter`](https://quay.io/repository/ricardbejarano/dnsmasq_exporter):

- [`1.4.6`, `latest` *(Dockerfile)*](Dockerfile)

### Building the container image yourself

```bash
docker build -t dnsmasq_exporter .
docker run -it -p 9153:9153 dnsmasq_exporter
```

### Building the binary yourself

```bash
go build -o bin/ .
./bin/dnsmasq_exporter
```

### Integrating with Prometheus

Add the following to `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: "dnsmasq"
    static_configs:
      - targets: ["<DNSMASQ_EXPORTER_ADDRESS>:9153"]
```
