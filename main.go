package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/miekg/dns"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/usewormhol/env"
)

var (
	servers = strings.Split(env.String("DNSMASQ_SERVERS", "127.0.0.1:53", env.Optional), ",")
	address = env.String("EXPORTER_LISTEN_ADDR", "127.0.0.1:9153", env.Optional)

	gauges = map[string]*prometheus.GaugeVec{
		"cachesize.bind.": prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dnsmasq_cachesize",
			Help: "configured cache size",
		}, []string{"instance"}),
		"insertions.bind.": prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dnsmasq_insertions",
			Help: "number of cache insertions during runtime",
		}, []string{"instance"}),
		"evictions.bind.": prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dnsmasq_evictions",
			Help: "number of cache evictions during runtime",
		}, []string{"instance"}),
		"misses.bind.": prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dnsmasq_misses",
			Help: "number of cache misses during runtime",
		}, []string{"instance"}),
		"hits.bind.": prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dnsmasq_hits",
			Help: "number of cache hits during runtime",
		}, []string{"instance"}),
		"auth.bind.": prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dnsmasq_auth",
			Help: "number of authoritative zone queries during runtime",
		}, []string{"instance"}),
	}

	client dns.Client
)

func main() {
	for _, gauge := range gauges {
		prometheus.MustRegister(gauge)
	}

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		questions := &dns.Msg{
			MsgHdr: dns.MsgHdr{
				Id:               dns.Id(),
				RecursionDesired: true,
			},
			Question: []dns.Question{
				{"cachesize.bind.", dns.TypeTXT, dns.ClassCHAOS},
				{"insertions.bind.", dns.TypeTXT, dns.ClassCHAOS},
				{"evictions.bind.", dns.TypeTXT, dns.ClassCHAOS},
				{"misses.bind.", dns.TypeTXT, dns.ClassCHAOS},
				{"hits.bind.", dns.TypeTXT, dns.ClassCHAOS},
				{"auth.bind.", dns.TypeTXT, dns.ClassCHAOS},
			},
		}
		for _, server := range servers {
			answers, _, err := client.Exchange(questions, server)
			if err != nil {
				log.Printf("server (%v) query error: %v\n", server, err)
				continue
			}
			for _, answer := range answers.Answer {
				txt, ok := answer.(*dns.TXT)
				if !ok {
					log.Printf("server (%v) answer parsing error: invalid answer type\n", server)
					continue
				}
				gauge, ok := gauges[txt.Hdr.Name]
				if !ok {
					log.Printf("server (%v) answer parsing error: invalid answer name\n", server)
					continue
				}
				value, err := strconv.ParseFloat(txt.Txt[0], 64)
				if err != nil {
					log.Printf("server (%v) answer parsing error: invalid answer value\n", server)
					continue
				}
				gauge.With(prometheus.Labels{"instance": server}).Set(value)
			}
		}
		promhttp.Handler().ServeHTTP(w, r)
	})

	log.Printf("dnsmasq servers %v\n", servers)
	log.Printf("metrics at http://%v/metrics\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
