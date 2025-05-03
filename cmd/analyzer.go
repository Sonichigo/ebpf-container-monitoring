package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	packetsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "container_packets_total",
			Help: "Total number of packets seen per container.",
		},
		[]string{"container"},
	)
	totalBytes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "container_bytes_total",
			Help: "Total bytes per container.",
		},
		[]string{"container"},
	)
	avgRtt = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "container_rtt_avg_ms",
			Help: "Average RTT in ms per container.",
		},
		[]string{"container"},
	)
)

func init() {
	prometheus.MustRegister(packetsTotal)
	prometheus.MustRegister(totalBytes)
	prometheus.MustRegister(avgRtt)
}

func main() {
	// Dummy values for demonstration. Replace with actual BPF polling logic.
	packetsTotal.WithLabelValues("nginx").Set(100)
	totalBytes.WithLabelValues("nginx").Set(204800)
	avgRtt.WithLabelValues("nginx").Set(12.5)

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Serving metrics at :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}