package process

import (
	"server/server/model"
)

func ProcessClick(c *(model.Counter)) error {
	c.Lock()
	c.Click++
	c.Unlock()

	return nil
}
