package main

import (
	"github.com/roganartu/orbit"
)

func requestReplicator(p orbit.Processor, ids []uint64) {
	// TODO send the request to the other servers in the cluster
	p.SetReplicatorIndex(ids[len(ids)-1])
}
