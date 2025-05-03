package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	Packets = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "container_packets_total",
			Help: "Packets per container",
		}, []string{"container"},
	)
	Bytes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "container_bytes_total",
			Help: "Bytes per container",
		}, []string{"container"},
	)
	RTT = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "container_rtt_avg_ms",
			Help: "RTT per container",
		}, []string{"container"},
	)
)

func Init() {
	prometheus.MustRegister(Packets, Bytes, RTT)
}