package main

import (
	"encoding/json"

	"github.com/roganartu/orbit"
)

func requestUnmarshaller(p orbit.Processor, ids []uint64) {
	for _, id := range ids {
		elem := p.GetMessage(id)
		elem.SetUnmarshalled(nil)

		if elem == nil || elem.GetMarshalled() == nil {
			continue
		}

		if b, ok := elem.GetMarshalled().([]byte); ok {
			var press Press
			err := json.Unmarshal(b, &press)
			if err != nil {
				continue
			}
			elem.SetUnmarshalled(press)
		}
	}
	p.SetUnmarshallerIndex(ids[len(ids)-1])
}
