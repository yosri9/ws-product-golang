package model

import "sync"

type Counters struct {
	sync.Mutex
	View  int
	Click int
}

var (
	CurrentCounter       *Counters
	SportCounter         = Counters{}
	EntertainmentCounter = Counters{}
	BusinessCounter      = Counters{}
	EducationCounter     = Counters{}
	MapCounter           map[string]string

	Content = []string{"sports", "entertainment", "business", "education"}
)
