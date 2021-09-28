package process

import (
	"server/server/model"
)

func ProcessClick(c *(model.Counter), data string) error {
	c.Lock()
	c.Click++
	c.Unlock()

	return nil
}
