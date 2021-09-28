package process

import (
	"server/server/model"
	"strconv"
)

func ProcessView(c *(model.Counter), data string, key string) error {
	c.Lock()
	_, currentTimeIsInMap := model.MapCounter[key]

	if currentTimeIsInMap == false {
		c.View = 0
		c.Click = 0
	}
	c.View++
	//value will be like views: 100, clicks: 4
	value := " view: " + strconv.Itoa(c.View) + " click: " + strconv.Itoa(c.Click)
	model.MapCounter[key] = value

	c.Unlock()
	return nil
}
