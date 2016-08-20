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

		if m, ok := elem.GetMarshalled().(*Message); ok {
			if m == nil {
				continue
			}

			var press Press
			err := json.Unmarshal(m.Body, &press)
			if err != nil {
				continue
			}
			m.Press = &press
			elem.SetUnmarshalled(m)
		}
	}
	p.SetUnmarshallerIndex(ids[len(ids)-1])
}
