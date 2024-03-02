package utils

// // GetMetricValue returns the value of a prometheus.Collector (https://stackoverflow.com/a/58875389/13180763)
// func GetMetricValue(col prometheus.Collector) float64 {
// 	var total float64

// 	collect(col, func(m *dto.Metric) { //nolint:govet
// 		if h := m.GetHistogram(); h != nil {
// 			total += float64(h.GetSampleCount())
// 		} else {
// 			total += m.GetCounter().GetValue()
// 		}
// 	})

// 	return total
// }

// collect calls the function for each metric associated with the Collector.
