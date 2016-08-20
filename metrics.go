package main

import (
	"math/big"

	"github.com/roganartu/orbit"
)

type Stats struct {
	TotalRequests *big.Int `json:"total_request_count"`
}

var (
	metrics *orbit.Loop
	stats   *Stats
)

func init() {
	stats = &Stats{
		TotalRequests: (&big.Int{}).SetInt64(0),
	}
	metrics = orbit.New(
		BUFFER_SIZE,
		nil,              // receiver
		nil,              // journaler
		nil,              // replicator
		nil,              // unmarshaller
		metricAggregator) // executor
	metrics.Start()
}

func metricAggregator(p orbit.Processor, ids []uint64) {
	for _, id := range ids {
		elem := p.GetMessage(id)
		if _, ok := elem.GetUnmarshalled().(*Message); ok {
			stats.TotalRequests.Add(stats.TotalRequests, one)
		}
	}
	p.SetExecutorIndex(ids[len(ids)-1])
}
