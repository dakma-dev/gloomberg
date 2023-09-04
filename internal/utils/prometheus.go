package utils

import (
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

// GetMetricValue returns the value of a prometheus.Collector (https://stackoverflow.com/a/58875389/13180763)
func GetMetricValue(col prometheus.Collector) float64 {
	var total float64

	collect(col, func(m dto.Metric) { //nolint:govet
		if h := m.GetHistogram(); h != nil {
			total += float64(h.GetSampleCount())
		} else {
			total += m.GetCounter().GetValue()
		}
	})

	return total
}

// collect calls the function for each metric associated with the Collector.
func collect(col prometheus.Collector, do func(dto.Metric)) {
	c := make(chan prometheus.Metric)

	go func(c chan prometheus.Metric) {
		col.Collect(c)
		close(c)
	}(c)

	for x := range c { // eg range across distinct label vector values
		m := dto.Metric{}
		_ = x.Write(&m)
		do(m) //nolint:govet
	}
}
