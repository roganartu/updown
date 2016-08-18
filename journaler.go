package main

import (
	"github.com/roganartu/orbit"
)

func requestJournaler(p orbit.Processor, ids []uint64) {

	p.SetJournalerIndex(ids[len(ids)-1])
}
