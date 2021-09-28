package model

import (
	"sync"
)

type Counters struct {
	sync.Mutex
	View  int
	Click int
	Name  string
	Time  string
}

var (
	CurrentCounter       *Counters
	SportCounter         = Counters{Name: "sport"}
	EntertainmentCounter = Counters{Name: "entertainment"}
	BusinessCounter      = Counters{Name: "business"}
	EducationCounter     = Counters{Name: "education"}
	MapCounter           map[string]string

	//ToDo counterMap : key={Map[event] = time} value ={map{view}=viewNumber ,map{click}= clickNumber} ??

	Content = []string{"sports", "entertainment", "business", "education"}
)
