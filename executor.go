package main

import (
	"math/big"

	"github.com/roganartu/orbit"
)

var count, one *big.Int

func init() {
	count = (&big.Int{}).SetInt64(0)
	one = (&big.Int{}).SetInt64(1)
}

func requestExecutor(p orbit.Processor, ids []uint64) {
	for _, id := range ids {
		elem := p.GetMessage(id)

		if elem == nil || elem.GetMarshalled() == nil {
			continue
		}

		if m, ok := elem.GetUnmarshalled().(*Message); ok {
			if m.Press.Up {
				count.Add(count, one)
			} else {
				count.Sub(count, one)
			}
			metrics.Input(m)
		}
	}
	p.SetExecutorIndex(ids[len(ids)-1])
}
