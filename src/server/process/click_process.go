package process

import (
	"server/model"
)

func ProcessClick(c *(model.Counters), data string) error {
	c.Lock()
	c.Click++
	c.Unlock()

	return nil
}
