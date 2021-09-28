package model

import (
	"sync"
)

type Counter struct {
	ID   int
	Name string
	Time string
	sync.Mutex
	View  int
	Click int
}

var (
	CurrentCounter       *Counter
	SportCounter         = Counter{Name: "sport"}
	EntertainmentCounter = Counter{Name: "entertainment"}
	BusinessCounter      = Counter{Name: "business"}
	EducationCounter     = Counter{Name: "education"}
	MapCounter           map[string]string

	//ToDo counterMap : key={Map[event] = time} value ={map{view}=viewNumber ,map{click}= clickNumber} ??

	Content = []string{"sports", "entertainment", "business", "education"}
)
