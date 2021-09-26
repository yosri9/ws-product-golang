package process

import (
	"server/model"
	"strconv"
)

func ProcessView(c *(model.Counters), data string, key string) error {
	c.Lock()
	_, currentTimeIsInMap := model.MapCounter[key]

	if currentTimeIsInMap == false {
		c.View = 0
		c.Click = 0
	}
	c.View++
	value := " view: " + strconv.Itoa(c.View) + " click: " + strconv.Itoa(c.Click)
	model.MapCounter[key] = value

	c.Unlock()
	return nil
}
