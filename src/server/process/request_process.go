package process

import (
	"math/rand"
	"net/http"
	"time"
)

func ProcessRequest(r *http.Request) error {
	time.Sleep(time.Duration(rand.Int31n(50)) * time.Millisecond)
	return nil
}
