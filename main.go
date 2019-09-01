package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Showmax/go-fqdn"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var startTime time.Time

func exporter() {
	go func() {
		fqdn.Get()
		for {
			timer.Inc()
			time.Since(startTime)
		}
		}()
	}
	
	var (
		host = promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "golang",
			Name:      "fully_qualified_domain_name",
			Help:      fqdn.Get(),
		})
		
		timer = promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "golang",
			Name:      "host_time_epoch",
			Help:      "The hosts uptime epoch",
		})
	)
	
	func main() {
		http.Handle("/metrics", promhttp.Handler())
		
		startTime = time.Now()
		exporter()
		fmt.Println("FQDN: ", fqdn.Get())
		
	http.ListenAndServe(":9999", nil)
}
